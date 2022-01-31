package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/osmosis-labs/osmosis/x/intergamm/types"
    "github.com/osmosis-labs/osmosis/x/intergamm/keeper"
    keepertest "github.com/osmosis-labs/osmosis/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IntergammKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
