package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v2/modules/core/05-port/types"
	gammaddr "github.com/osmosis-labs/osmosis/v043_temp/address"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

	addr := gammaddr.Module("intergamm", []byte(fmt.Sprintf("%s/%s", req.Port, req.Channel)))

	return &types.QueryChannelAddressResponse{
		Address: sdk.AccAddress(addr).String(),
	}, nil
}
