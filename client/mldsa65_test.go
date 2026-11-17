package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	cometcrypto "github.com/cometbft/cometbft/crypto"
	ced25519 "github.com/cometbft/cometbft/crypto/ed25519"
	cryptoenc "github.com/cometbft/cometbft/crypto/encoding"
	cmldsa65 "github.com/cometbft/cometbft/crypto/mldsa65"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	slclient "github.com/strangelove-ventures/cometbft-client/client"
	slpc "github.com/strangelove-ventures/cometbft-client/proto/tendermint/crypto"
	"github.com/stretchr/testify/suite"
)

type PubKeyDecodeSuite struct {
	suite.Suite

	// results maps a JSON-RPC method to the result the fake node returns for it.
	results map[string]any
	raw     *slclient.Client
	client  RPCClient
}

type oneofSum interface {
	Size() int
	MarshalTo([]byte) (int, error)
}

// encode checks both marshal paths agree, since each writes the key independently.
func (s *PubKeyDecodeSuite) encode(sum oneofSum) []byte {
	b := make([]byte, sum.Size())
	n, err := sum.MarshalTo(b)
	s.Require().NoError(err)

	sizedSum, ok := sum.(interface {
		MarshalToSizedBuffer([]byte) (int, error)
	})
	s.Require().True(ok)
	sized := make([]byte, sum.Size())
	m, err := sizedSum.MarshalToSizedBuffer(sized)
	s.Require().NoError(err)
	s.Require().Equal(b[:n], sized[len(sized)-m:])
	return b[:n]
}

func TestPubKeyDecodeSuite(t *testing.T) {
	suite.Run(t, new(PubKeyDecodeSuite))
}

func (s *PubKeyDecodeSuite) SetupTest() {
	s.results = map[string]any{}

	// Results are encoded with cometbft's own JSON codec, so the test sees what a real node sends.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
		}
		// Require would call FailNow, which is only valid on the test goroutine.
		s.NoError(json.NewDecoder(r.Body).Decode(&req))

		result, err := cmtjson.Marshal(s.results[req.Method])
		s.NoError(err)

		s.NoError(json.NewEncoder(w).Encode(map[string]any{
			"jsonrpc": "2.0",
			"id":      req.ID,
			"result":  json.RawMessage(result),
		}))
	}))
	s.T().Cleanup(server.Close)

	c, err := slclient.NewClient(server.URL, 5*time.Second)
	s.Require().NoError(err)
	s.raw = c
	s.client = NewRPCClient(c)
}

func (s *PubKeyDecodeSuite) TestProtoMldsa65PubKeyCompare() {
	key := &protoMldsa65PubKey{Mldsa65: []byte{1, 2, 3}}

	testCases := []struct {
		name    string
		other   interface{}
		equal   bool
		compare int
	}{
		{"same key", &protoMldsa65PubKey{Mldsa65: []byte{1, 2, 3}}, true, 0},
		{"same key by value", protoMldsa65PubKey{Mldsa65: []byte{1, 2, 3}}, true, 0},
		{"smaller key", &protoMldsa65PubKey{Mldsa65: []byte{1, 2, 2}}, false, 1},
		{"larger key", &protoMldsa65PubKey{Mldsa65: []byte{1, 2, 4}}, false, -1},
		{"empty ed25519 key", &slpc.PublicKey_Ed25519{}, false, 1}, // the embedded ed25519 methods report equal
		{"nil", nil, false, 1},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Require().Equal(tc.equal, key.Equal(tc.other))
			s.Require().Equal(tc.compare, key.Compare(tc.other))
		})
	}
}

func (s *PubKeyDecodeSuite) TestMldsa65Adapter() {
	priv, err := cmldsa65.GenPrivKey()
	s.Require().NoError(err)
	otherPriv, err := cmldsa65.GenPrivKey()
	s.Require().NoError(err)

	pubKey := priv.PubKey().(cmldsa65.PubKey)
	adapter := mldsa65PubKey{pubKey}

	s.Require().Equal([]byte(pubKey.Address()), []byte(adapter.Address()))
	s.Require().True(adapter.Equals(mldsa65PubKey{pubKey}))
	s.Require().False(adapter.Equals(mldsa65PubKey{otherPriv.PubKey().(cmldsa65.PubKey)}))
	s.Require().Equal(cometcrypto.PubKey(pubKey), convertPubKey(adapter))
}

func (s *PubKeyDecodeSuite) TestDecodeValidatorKeys() {
	mldsaPriv, err := cmldsa65.GenPrivKey()
	s.Require().NoError(err)

	testCases := []struct {
		name   string
		pubKey cometcrypto.PubKey
	}{
		{"ed25519", ced25519.GenPrivKey().PubKey()},
		{"ml-dsa-65", mldsaPriv.PubKey()},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctx := context.Background()
			height := int64(5)
			validator := tmtypes.NewValidator(tc.pubKey, 10)

			s.results["status"] = &coretypes.ResultStatus{
				ValidatorInfo: coretypes.ValidatorInfo{
					Address:     tc.pubKey.Address(),
					PubKey:      tc.pubKey,
					VotingPower: 10,
				},
			}
			s.results["validators"] = &coretypes.ResultValidators{
				BlockHeight: height,
				Validators:  []*tmtypes.Validator{validator},
				Count:       1,
				Total:       1,
			}
			protoPubKey, err := cryptoenc.PubKeyToProto(tc.pubKey)
			s.Require().NoError(err)
			s.results["block_results"] = &coretypes.ResultBlockResults{
				Height:           height,
				ValidatorUpdates: []abci.ValidatorUpdate{{PubKey: protoPubKey, Power: 10}},
			}

			status, err := s.client.Status(ctx)
			s.Require().NoError(err)
			s.Require().Equal(tc.pubKey, status.ValidatorInfo.PubKey)

			vals, err := s.client.Validators(ctx, &height, nil, nil)
			s.Require().NoError(err)
			s.Require().Len(vals.Validators, 1)
			s.Require().Equal(tc.pubKey, vals.Validators[0].PubKey)
			s.Require().Equal(validator.Address, vals.Validators[0].Address)

			// The IBC header carries the validator set proto-encoded, with the hash the light client checks.
			valSet := tmtypes.NewValidatorSet(vals.Validators)
			s.Require().Equal(tmtypes.NewValidatorSet([]*tmtypes.Validator{validator}).Hash(), valSet.Hash())
			_, err = valSet.ToProto()
			s.Require().NoError(err)

			raw, err := s.raw.BlockResults(ctx, &height)
			s.Require().NoError(err)
			s.Require().Len(raw.ValidatorUpdates, 1)
			decoded := raw.ValidatorUpdates[0].PubKey.Sum
			s.Require().Equal(s.encode(protoPubKey.Sum), s.encode(decoded))

			res, err := s.client.BlockResults(ctx, &height)
			s.Require().NoError(err)
			s.Require().Empty(res.ValidatorUpdates)
		})
	}
}
