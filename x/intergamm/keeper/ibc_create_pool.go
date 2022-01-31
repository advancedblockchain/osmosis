package keeper

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v2/modules/core/24-host"
	gammaddr "github.com/osmosis-labs/osmosis/v043_temp/address"
	gammbalancer "github.com/osmosis-labs/osmosis/x/gamm/pool-models/balancer"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
)

// TransmitIbcCreatePoolPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcCreatePoolPacket(
	ctx sdk.Context,
	packetData types.IbcCreatePoolPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.ChannelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvIbcCreatePoolPacket processes packet reception
func (k Keeper) OnRecvIbcCreatePoolPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcCreatePoolPacketData) (packetAck types.IbcCreatePoolPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// Convert the intergamm types to original gamm types
	sender := gammaddr.Module(types.ModuleName, []byte(fmt.Sprintf("%s/%s", packet.SourcePort, packet.SourceChannel)))
	var changeParams *gammbalancer.SmoothWeightChangeParams
	if data.Params != nil {
		changeParams = &gammbalancer.SmoothWeightChangeParams{
			StartTime: data.Params.SmoothWeightChangeParams.StartTime,
			Duration:  data.Params.SmoothWeightChangeParams.Duration,
		}

		changeParams.InitialPoolWeights = make([]gammtypes.PoolAsset, len(data.Params.SmoothWeightChangeParams.InitialPoolWeights))
		for i := 0; i < len(data.Assets); i++ {
			changeParams.InitialPoolWeights[i] = gammtypes.PoolAsset(data.Params.SmoothWeightChangeParams.InitialPoolWeights[i])
		}
		changeParams.TargetPoolWeights = make([]gammtypes.PoolAsset, len(data.Params.SmoothWeightChangeParams.TargetPoolWeights))
		for i := 0; i < len(data.Assets); i++ {
			changeParams.TargetPoolWeights[i] = gammtypes.PoolAsset(data.Params.SmoothWeightChangeParams.TargetPoolWeights[i])
		}
	}
	params := gammbalancer.BalancerPoolParams{
		SwapFee:                  data.Params.SwapFee,
		ExitFee:                  data.Params.ExitFee,
		SmoothWeightChangeParams: changeParams,
	}
	assets := make([]gammtypes.PoolAsset, len(data.Assets))
	for i := 0; i < len(data.Assets); i++ {
		assets[i] = gammtypes.PoolAsset(*data.Assets[i])
	}

	poolId, err := k.gammKeeper.CreateBalancerPool(ctx, sender, params, assets, data.FuturePoolGovernor)
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
