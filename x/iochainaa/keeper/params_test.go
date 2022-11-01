package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochainaa/testutil/keeper"
	"github.com/mytestlab123/iochainaa/x/iochainaa/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IochainaaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
