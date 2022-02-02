package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	gammaddr "github.com/osmosis-labs/osmosis/v043_temp/address"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
)

// OnRecvIbcCreatePoolPacket processes packet reception
func (k Keeper) OnRecvIbcCreatePoolPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcCreatePoolPacketData) (packetAck types.IbcCreatePoolPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	sender := gammaddr.Module(types.ModuleName, []byte(fmt.Sprintf("%s/%s", packet.SourcePort, packet.SourceChannel)))

	poolId, err := k.gammKeeper.CreateBalancerPool(ctx, sender, *data.Params, data.Assets, data.FuturePoolGovernor)
	if err != nil {
		return packetAck, err
	}
	packetAck.PoolId = poolId

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			gammtypes.TypeEvtPoolCreated,
			sdk.NewAttribute(gammtypes.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, gammtypes.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, sdk.AccAddress(sender).String()),
		),
	})

	return packetAck, nil
}
