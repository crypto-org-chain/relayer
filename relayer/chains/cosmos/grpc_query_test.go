package cosmos_test

import (
	"bytes"
	"strings"
	"testing"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/relayer/v2/relayer/chains/cosmos"
	"github.com/stretchr/testify/suite"
)

type GRPCCodecSuite struct {
	suite.Suite
}

func TestGRPCCodecSuite(t *testing.T) {
	suite.Run(t, new(GRPCCodecSuite))
}

func (s *GRPCCodecSuite) TestRoundTrip() {
	testCases := []struct {
		name string
		msg  banktypes.QueryBalanceRequest
	}{
		{"empty", banktypes.QueryBalanceRequest{}},
		{"small", banktypes.QueryBalanceRequest{Address: "cosmos1abc", Denom: "stake"}},
		// large enough for grpc to marshal into a pooled buffer
		{"pooled", banktypes.QueryBalanceRequest{Denom: strings.Repeat("a", 1<<16)}},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			bz, err := cosmos.GRPCMarshal(&tc.msg)
			s.Require().NoError(err)

			expBz, err := tc.msg.Marshal()
			s.Require().NoError(err)
			// bytes.Equal treats the empty message's nil and zero-length encodings as the same.
			s.Require().True(bytes.Equal(expBz, bz))

			var decoded banktypes.QueryBalanceRequest
			s.Require().NoError(cosmos.GRPCUnmarshal(bz, &decoded))
			s.Require().Equal(tc.msg, decoded)
		})
	}
}

func (s *GRPCCodecSuite) TestUnmarshalInvalid() {
	var decoded banktypes.QueryBalanceRequest
	s.Require().Error(cosmos.GRPCUnmarshal([]byte{0xff}, &decoded))
}
