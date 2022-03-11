package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
)

// OnRecvIbcExitPoolPacket processes packet reception
func (k Keeper) OnRecvIbcExitPoolPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcExitPoolPacketData) (packetAck types.IbcExitPoolPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	sender := genChannelAddress(packet.SourcePort, packet.SourceChannel)

	err = k.gammKeeper.ExitPool(ctx, sender, data.PoolId, data.ShareInAmount, data.TokenOutMins)
	if err != nil {
		return packetAck, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			gammtypes.TypeEvtPoolExited,
			sdk.NewAttribute(gammtypes.AttributeKeyPoolId, strconv.FormatUint(data.PoolId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, gammtypes.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, sdk.AccAddress(sender).String()),
		),
	})

	return packetAck, nil
}
