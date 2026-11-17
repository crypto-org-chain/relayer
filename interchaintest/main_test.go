package interchaintest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"go.uber.org/zap"
)

const configuredChainsEnv = "IBCTEST_CONFIGURED_CHAINS"

// builtinGenesisPatches holds genesis changes for built-in chains whose public images differ from the ones these tests were written for.
var builtinGenesisPatches = map[string][]cosmos.GenesisKV{
	// Gaia v19+ ships x/feemarket with fee denom "stake" and a minimum gas price of 1uatom, far above the
	// configured gas price. Its fallback fee check keeps that minimum even with feemarket disabled, so lower it to
	// the mainnet value instead. Gaia v19+ also rejects zero-fee txs either way.
	"gaia": {
		cosmos.NewGenesisKV("app_state.feemarket.params.fee_denom", "uatom"),
		cosmos.NewGenesisKV("app_state.feemarket.params.min_base_gas_price", "0.005"),
		cosmos.NewGenesisKV("app_state.feemarket.state.base_gas_price", "0.005"),
	},
}

// TestMain points interchaintest at chains.yaml unless the caller already set its own chain configs.
func TestMain(m *testing.M) {
	if os.Getenv(configuredChainsEnv) == "" {
		path, err := filepath.Abs("chains.yaml")
		if err == nil {
			err = os.Setenv(configuredChainsEnv, path)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "set %s: %v\n", configuredChainsEnv, err)
			os.Exit(1)
		}
	}
	os.Exit(m.Run())
}

// NewChainFactory builds a built-in chain factory, applying builtinGenesisPatches to each spec that does not set its own ModifyGenesis.
func NewChainFactory(log *zap.Logger, specs []*interchaintest.ChainSpec) *interchaintest.BuiltinChainFactory {
	for _, s := range specs {
		if kv, ok := builtinGenesisPatches[s.Name]; ok && s.ModifyGenesis == nil {
			s.ModifyGenesis = cosmos.ModifyGenesis(kv)
		}
	}
	return interchaintest.NewBuiltinChainFactory(log, specs)
}
