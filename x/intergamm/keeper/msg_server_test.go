package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/osmosis-labs/osmosis/testutil/keeper"
	"github.com/osmosis-labs/osmosis/x/intergamm/keeper"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IntergammKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
