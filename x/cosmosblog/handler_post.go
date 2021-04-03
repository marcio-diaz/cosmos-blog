// x/blog/handler_post.go
package cosmosblog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/marcio.diaz/cosmos-blog/x/cosmosblog/keeper"
	"github.com/marcio.diaz/cosmos-blog/x/cosmosblog/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
