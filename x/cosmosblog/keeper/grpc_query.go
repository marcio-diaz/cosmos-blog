package keeper

import (
	"github.com/marcio.diaz/cosmos-blog/x/cosmosblog/types"
)

var _ types.QueryServer = Keeper{}
