package airesult

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oraichain/orai/x/airesult/types"
	provider "github.com/oraichain/orai/x/provider/types"
)

// InitGenesis initialize default parameters, can assign some coins in the chain here
// and the k's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	// TODO: Define logic for when you would like to initialize a new genesis
	// Init params for the airequest module
	k.SetParam(ctx, provider.KeyOracleScriptRewardPercentage, data.Params.OracleScriptRewardPercentage)
	k.SetParam(ctx, types.KeyExpirationCount, data.Params.ExpirationCount)
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	// TODO: Define logic for exporting state
	return types.GenesisState{
		AIRequestResults: []types.AIRequestResult{},
		Params:           k.GetParams(ctx),
	}
}