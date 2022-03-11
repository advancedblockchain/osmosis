package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v2/modules/core/05-port/types"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ChannelAddress returns the channel address associated with the given port and channel-id. Note that it will return error if there's no channel open
// with the given port and channel-id.
func (k Keeper) ChannelAddress(goCtx context.Context, req *types.QueryChannelAddressRequest) (*types.QueryChannelAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	boundPort := k.GetPort(ctx)
	if boundPort != req.Port {
		return nil, sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", req.Port, boundPort)
	}

	_, found := k.ChannelKeeper.GetChannel(ctx, req.Port, req.Channel)
	if !found {
		return nil, sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", req.Port, req.Channel)
	}

	addr := genChannelAddress(req.Port, req.Channel)

	return &types.QueryChannelAddressResponse{
		Address: sdk.AccAddress(addr).String(),
	}, nil
}
