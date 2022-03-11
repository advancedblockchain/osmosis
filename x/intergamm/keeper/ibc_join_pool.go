package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
)

// OnRecvIbcJoinPoolPacket processes packet reception
func (k Keeper) OnRecvIbcJoinPoolPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcJoinPoolPacketData) (packetAck types.IbcJoinPoolPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	sender := genChannelAddress(packet.SourcePort, packet.SourceChannel)

	err = k.gammKeeper.JoinPool(ctx, sender, data.PoolId, data.ShareOutAmount, data.TokenInMaxs)
	if err != nil {
		return packetAck, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			gammtypes.TypeEvtPoolJoined,
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
