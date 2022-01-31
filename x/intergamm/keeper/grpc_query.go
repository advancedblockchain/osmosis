package keeper

import (
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
)

var _ types.QueryServer = Keeper{}
