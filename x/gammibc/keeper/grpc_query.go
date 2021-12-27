package keeper

import (
	"github.com/osmosis-labs/osmosis/x/gammibc/types"
)

var _ types.QueryServer = Keeper{}
