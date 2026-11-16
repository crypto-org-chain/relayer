package cosmos

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stretchr/testify/suite"
)

var (
	coinType = 118

	testCfg = CosmosProviderConfig{
		KeyDirectory:     "",
		Key:              "default",
		ChainName:        "osmosis",
		ChainID:          "osmosis-1",
		AccountPrefix:    "osmo",
		KeyringBackend:   "test",
		DynamicGasPrice:  true,
		GasAdjustment:    1.2,
		GasPrices:        "0.0025uosmo",
		MinGasAmount:     1,
		MaxGasAmount:     0,
		Debug:            false,
		Timeout:          "30s",
		BlockTimeout:     "30s",
		OutputFormat:     "json",
		SignModeStr:      "direct",
		ExtraCodecs:      nil,
		Modules:          nil,
		Slip44:           &coinType,
		SigningAlgorithm: "",
		Broadcast:        "batch",
		MinLoopDuration:  0,
		ExtensionOptions: nil,
		FeeGrants:        nil,
	}
)

type FeeMarketTestSuite struct {
	suite.Suite
}

func TestFeeMarketTestSuite(t *testing.T) {
	suite.Run(t, new(FeeMarketTestSuite))
}

// newABCIQueryServer answers the base fee abci_query with resp and rejects any other request.
func (s *FeeMarketTestSuite) newABCIQueryServer(resp abci.ResponseQuery) *httptest.Server {
	result, err := cmtjson.Marshal(coretypes.ResultABCIQuery{Response: resp})
	s.Require().NoError(err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
			Params struct {
				Path string `json:"path"`
			} `json:"params"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Method != "abci_query" || req.Params.Path != queryPath {
			http.Error(w, "unexpected request: "+req.Method+" "+req.Params.Path, http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"jsonrpc": "2.0",
			"id":      req.ID,
			"result":  json.RawMessage(result),
		})
	}))
	s.T().Cleanup(srv.Close)
	return srv
}

func (s *FeeMarketTestSuite) TestQueryBaseFee() {
	testCases := []struct {
		name      string
		resp      abci.ResponseQuery
		expFee    string
		expErrMsg string
	}{
		{
			name: "base fee with 18 implied decimals",
			// Osmosis returns QueryEipBaseFeeResponse: field 1 tag (0x0a) and length (0x10) before the integer Dec string.
			resp:   abci.ResponseQuery{Value: []byte("\n\x102500000000000000")},
			expFee: "0.002500uosmo",
		},
		{
			name:   "non-zero response code",
			resp:   abci.ResponseQuery{Code: 1},
			expFee: "",
		},
		{
			name:      "unparsable base fee",
			resp:      abci.ResponseQuery{Value: []byte("\n\x05abcde")},
			expErrMsg: "failed to set decimal string",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cfg := testCfg
			cfg.RPCAddr = s.newABCIQueryServer(tc.resp).URL

			p, err := cfg.NewProvider(nil, s.T().TempDir(), true, cfg.ChainName)
			s.Require().NoError(err)

			ctx := context.Background()
			s.Require().NoError(p.Init(ctx))

			baseFee, err := p.(*CosmosProvider).QueryBaseFee(ctx)
			if tc.expErrMsg != "" {
				s.Require().ErrorContains(err, tc.expErrMsg)
				return
			}
			s.Require().NoError(err)
			s.Require().Equal(tc.expFee, baseFee)
		})
	}
}

func (s *FeeMarketTestSuite) TestParseDenom() {
	denom, err := parseTokenDenom(testCfg.GasPrices)
	s.Require().NoError(err)
	s.Require().Equal("uosmo", denom)
}
