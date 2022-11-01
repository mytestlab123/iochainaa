package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mytestlab123/iochainaa/testutil/keeper"
	"github.com/mytestlab123/iochainaa/x/iochainaa/keeper"
	"github.com/mytestlab123/iochainaa/x/iochainaa/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IochainaaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
