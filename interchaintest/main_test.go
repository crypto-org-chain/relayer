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

// Genesis changes for built-in chains whose public images differ from the ones these tests were written for.
// Keyed by spec Name and applied to every version, so the gaia entry assumes gaia v19+.
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

// Keeps a caller-set chain config so local runs can try other images.
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

// A spec with its own ModifyGenesis skips the built-in patch, so it must include those changes itself.
func NewChainFactory(log *zap.Logger, specs []*interchaintest.ChainSpec) *interchaintest.BuiltinChainFactory {
	for _, s := range specs {
		if kv, ok := builtinGenesisPatches[s.Name]; ok && s.ModifyGenesis == nil {
			s.ModifyGenesis = cosmos.ModifyGenesis(kv)
		}
	}
	return interchaintest.NewBuiltinChainFactory(log, specs)
}
