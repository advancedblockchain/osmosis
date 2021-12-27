package keeper

import (
	"context"

	"github.com/osmosis-labs/osmosis/x/gammibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
)

func (k msgServer) SendIbcCreatePool(goCtx context.Context, msg *types.MsgSendIbcCreatePool) (*types.MsgSendIbcCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcCreatePoolPacketData

	packet.Weights = msg.Weights
	packet.InitialDeposit = msg.InitialDeposit
	packet.SwapFee = msg.SwapFee
	packet.ExitFee = msg.ExitFee
	packet.FutureGovernor = msg.FutureGovernor

	// Transmit the packet
	err := k.TransmitIbcCreatePoolPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcCreatePoolResponse{}, nil
}
