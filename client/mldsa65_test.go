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
	"github.com/stretchr/testify/suite"
)

type PubKeyDecodeSuite struct {
	suite.Suite

	// results maps a JSON-RPC method to the result the fake node returns for it.
	results map[string]any
	client  RPCClient
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
		s.Require().NoError(json.NewDecoder(r.Body).Decode(&req))

		result, err := cmtjson.Marshal(s.results[req.Method])
		s.Require().NoError(err)

		s.Require().NoError(json.NewEncoder(w).Encode(map[string]any{
			"jsonrpc": "2.0",
			"id":      req.ID,
			"result":  json.RawMessage(result),
		}))
	}))
	s.T().Cleanup(server.Close)

	c, err := slclient.NewClient(server.URL, 5*time.Second)
	s.Require().NoError(err)
	s.client = NewRPCClient(c)
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

			_, err = s.client.BlockResults(ctx, &height)
			s.Require().NoError(err)
		})
	}
}
