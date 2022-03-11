package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
)

// OnRecvIbcWithdrawPacket processes packet reception
func (k Keeper) OnRecvIbcWithdrawPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcWithdrawPacketData) (packetAck types.IbcWithdrawPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	sender := genChannelAddress(packet.SourcePort, packet.SourceChannel)

	// TODO: Investigate the relayer sequence mismatch error when we withdraw more than one asset
	for _, as := range data.Assets {
		err = k.tansferKeeper.SendTransfer(ctx, data.TransferPort, data.TransferChannel, as, sender, data.Receiver, clienttypes.NewHeight(0, 1000), 0) // TODO: Think about better values for timeout height and timestamp or get them from ibc packet
		if err != nil {
			return packetAck, err
		}
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIbcWithdrawPacket,
			sdk.NewAttribute(types.AttributeKeyAssets, sdk.Coins(data.Assets).String()),
		),
	})

	return packetAck, nil
}
