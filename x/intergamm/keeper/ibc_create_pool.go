package keeper

import (
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
