package tests

import (
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
)

var (
	image = ibc.DockerImage{
		Repository: "oraichain/foundation-orai",
		Version:    "0.41.3-alpine-dev",
		UidGid:     "1025:1025",
	}

	chainConfig = ibc.ChainConfig{
		Type:                "cosmos",
		Name:                "Oraichain",
		ChainID:             "Oraichain-e2e",
		Images:              []ibc.DockerImage{image},
		Bin:                 "oraid",
		Bech32Prefix:        "orai",
		Denom:               "orai",
		CoinType:            "118",
		GasPrices:           "0.01orai",
		GasAdjustment:       1.1,
		TrustingPeriod:      "112h",
		NoHostMount:         false,
		SkipGenTx:           false,
		PreGenesis:          nil,
		ModifyGenesis:       nil,
		ConfigFileOverrides: nil,
		EncodingConfig:      chainEncoding(),
	}

	// pathJunoGaia        = "juno-gaia"
	genesisWalletAmount = int64(10_000_000)
)

// chainEncoding registers the Juno specific module codecs so that the associated types and msgs
func chainEncoding() *simappparams.EncodingConfig {
	cfg := cosmos.DefaultEncoding()
	return &cfg
}
