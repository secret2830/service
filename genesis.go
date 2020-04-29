package service

import (
	"encoding/hex"
	"fmt"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis - store genesis parameters
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	k.SetParams(ctx, data.Params)

	for _, definition := range data.Definitions {
		k.SetServiceDefinition(ctx, definition)
	}

	for _, binding := range data.Bindings {
		k.SetServiceBinding(ctx, binding)
		k.SetOwnerServiceBinding(ctx, binding)
	}

	for ownerAddressStr, withdrawAddress := range data.WithdrawAddresses {
		ownerAddress, _ := sdk.AccAddressFromBech32(ownerAddressStr)
		k.SetWithdrawAddress(ctx, ownerAddress, withdrawAddress)
	}

	for reqContextIDStr, requestContext := range data.RequestContexts {
		requestContextID, _ := hex.DecodeString(reqContextIDStr)
		k.SetRequestContext(ctx, requestContextID, requestContext)
	}
}

// ExportGenesis - output genesis parameters
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	definitions := []ServiceDefinition{}
	bindings := []ServiceBinding{}
	withdrawAddresses := make(map[string]sdk.AccAddress)
	requestContexts := make(map[string]RequestContext)

	k.IterateServiceDefinitions(
		ctx,
		func(definition ServiceDefinition) bool {
			definitions = append(definitions, definition)
			return false
		},
	)

	k.IterateServiceBindings(
		ctx,
		func(binding ServiceBinding) bool {
			bindings = append(bindings, binding)
			return false
		},
	)

	k.IterateWithdrawAddresses(
		ctx,
		func(ownerAddress sdk.AccAddress, withdrawAddress sdk.AccAddress) bool {
			withdrawAddresses[ownerAddress.String()] = withdrawAddress
			return false
		},
	)

	k.IterateRequestContexts(
		ctx,
		func(requestContextID tmbytes.HexBytes, requestContext RequestContext) bool {
			requestContexts[requestContextID.String()] = requestContext
			return false
		},
	)

	return NewGenesisState(
		k.GetParams(ctx),
		definitions,
		bindings,
		withdrawAddresses,
		requestContexts,
	)
}

// PrepForZeroHeightGenesis refunds the deposits, service fees and earned fees
func PrepForZeroHeightGenesis(ctx sdk.Context, k Keeper) {
	// refund service fees from all active requests
	if err := k.RefundServiceFees(ctx); err != nil {
		panic(fmt.Sprintf("failed to refund the service fees: %s", err))
	}

	// refund all the earned fees
	if err := k.RefundEarnedFees(ctx); err != nil {
		panic(fmt.Sprintf("failed to refund the earned fees: %s", err))
	}

	// reset request contexts state and batch
	if err := k.ResetRequestContextsStateAndBatch(ctx); err != nil {
		panic(fmt.Sprintf("failed to reset the request context state: %s", err))
	}
}
