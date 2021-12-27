package gammibc_test

import (
	"testing"

	keepertest "github.com/osmosis-labs/osmosis/testutil/keeper"
	"github.com/osmosis-labs/osmosis/testutil/nullify"
	"github.com/osmosis-labs/osmosis/x/gammibc"
	"github.com/osmosis-labs/osmosis/x/gammibc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GammibcKeeper(t)
	gammibc.InitGenesis(ctx, *k, genesisState)
	got := gammibc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
