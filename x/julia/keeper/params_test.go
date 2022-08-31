package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "julia/testutil/keeper"
	"julia/x/julia/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JuliaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
