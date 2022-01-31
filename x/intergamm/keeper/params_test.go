package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/osmosis-labs/osmosis/testutil/keeper"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IntergammKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
