package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	gammaddr "github.com/osmosis-labs/osmosis/v043_temp/address"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
	"github.com/tendermint/starport/starport/pkg/cosmosibckeeper"
)

type (
	Keeper struct {
		*cosmosibckeeper.Keeper
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		gammKeeper    types.GammKeeper
		tansferKeeper types.TransferKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper cosmosibckeeper.ChannelKeeper,
	portKeeper cosmosibckeeper.PortKeeper,
	scopedKeeper cosmosibckeeper.ScopedKeeper,
	gammKeeper types.GammKeeper,
	tansferKeeper types.TransferKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		Keeper: cosmosibckeeper.NewKeeper(
			types.PortKey,
			storeKey,
			channelKeeper,
			portKeeper,
			scopedKeeper,
		),
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		gammKeeper:    gammKeeper,
		tansferKeeper: tansferKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Generates an address for port and channel-id pair.
// NOTE: Always use source port and channel-id to get an unique address
func genChannelAddress(srcPort, srcChannel string) []byte {
	// TODO: Use standard sdk methods for this
	return gammaddr.Module(types.ModuleName, []byte(fmt.Sprintf("%s/%s", srcPort, srcChannel)))
}
