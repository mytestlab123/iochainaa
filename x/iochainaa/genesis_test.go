package iochainaa_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochainaa/testutil/keeper"
	"github.com/mytestlab123/iochainaa/testutil/nullify"
	"github.com/mytestlab123/iochainaa/x/iochainaa"
	"github.com/mytestlab123/iochainaa/x/iochainaa/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IochainaaKeeper(t)
	iochainaa.InitGenesis(ctx, *k, genesisState)
	got := iochainaa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
