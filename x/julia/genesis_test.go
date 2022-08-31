package julia_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "julia/testutil/keeper"
	"julia/testutil/nullify"
	"julia/x/julia"
	"julia/x/julia/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JuliaKeeper(t)
	julia.InitGenesis(ctx, *k, genesisState)
	got := julia.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
