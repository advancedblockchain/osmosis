package keeper_test

import (
	"testing"

	testkeeper "github.com/osmosis-labs/osmosis/testutil/keeper"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IntergammKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
