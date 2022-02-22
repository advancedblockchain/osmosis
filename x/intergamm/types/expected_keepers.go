package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	gammbalancer "github.com/osmosis-labs/osmosis/x/gamm/pool-models/balancer"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type GammKeeper interface {
	CreateBalancerPool(
		ctx sdk.Context,
		sender sdk.AccAddress,
		BalancerPoolParams gammbalancer.BalancerPoolParams,
		poolAssets []gammtypes.PoolAsset,
		futurePoolGovernor string,
	) (uint64, error)

	JoinPool(
		ctx sdk.Context,
		sender sdk.AccAddress,
		poolId uint64,
		shareOutAmount sdk.Int,
		tokenInMaxs sdk.Coins,
	) error

	ExitPool(
		ctx sdk.Context,
		sender sdk.AccAddress,
		poolId uint64,
		shareInAmount sdk.Int,
		tokenOutMins sdk.Coins,
	) error
}

type TransferKeeper interface {
	SendTransfer(
		ctx sdk.Context,
		sourcePort,
		sourceChannel string,
		token sdk.Coin,
		sender sdk.AccAddress,
		receiver string,
		timeoutHeight clienttypes.Height,
		timeoutTimestamp uint64,
	) error
}
