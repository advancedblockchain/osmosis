package intergamm_test

import (
	"testing"

	keepertest "github.com/osmosis-labs/osmosis/testutil/keeper"
	"github.com/osmosis-labs/osmosis/testutil/nullify"
	"github.com/osmosis-labs/osmosis/x/intergamm"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IntergammKeeper(t)
	intergamm.InitGenesis(ctx, *k, genesisState)
	got := intergamm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
