package keeper

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	gammaddr "github.com/osmosis-labs/osmosis/v043_temp/address"
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

	// TODO: emit events related to creating a pool

	return packetAck, nil
}

// OnAcknowledgementIbcCreatePoolPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcCreatePoolPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcCreatePoolPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcCreatePoolPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIbcCreatePoolPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcCreatePoolPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcCreatePoolPacketData) error {

	// TODO: packet timeout logic

	return nil
}
