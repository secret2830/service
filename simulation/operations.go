package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irismod/service/keeper"
	"github.com/irismod/service/simapp/helpers"
	simappparams "github.com/irismod/service/simapp/params"
	"github.com/irismod/service/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgDefineService         = "op_weight_msg_define_service"
	OpWeightMsgBindService           = "op_weight_msg_bind_service"
	OpWeightMsgUpdateServiceBinding  = "op_weight_msg_update_service_binding"
	OpWeightMsgSetWithdrawAddress    = "op_weight_msg_set_withdraw_address"
	OpWeightMsgDisableServiceBinding = "op_weight_msg_disable_service_binding"
	OpWeightMsgEnableServiceBinding  = "op_weight_msg_enable_service_binding"
	OpWeightMsgRefundServiceDeposit  = "op_weight_msg_refund_service_deposit"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simulation.AppParams,
	cdc *codec.Codec,
	ak types.AccountKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgDefineService         int
		weightMsgBindService           int
		weightMsgUpdateServiceBinding  int
		weightMsgSetWithdrawAddress    int
		weightMsgDisableServiceBinding int
		weightMsgEnableServiceBinding  int
		weightMsgRefundServiceDeposit  int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgDefineService, &weightMsgDefineService, nil,
		func(_ *rand.Rand) {
			weightMsgDefineService = simappparams.DefaultWeightMsgDefineService
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBindService, &weightMsgBindService, nil,
		func(_ *rand.Rand) {
			weightMsgBindService = simappparams.DefaultWeightMsgBindService
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateServiceBinding, &weightMsgUpdateServiceBinding, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateServiceBinding = simappparams.DefaultWeightMsgUpdateServiceBinding
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSetWithdrawAddress, &weightMsgSetWithdrawAddress, nil,
		func(_ *rand.Rand) {
			weightMsgSetWithdrawAddress = simappparams.DefaultWeightMsgSetWithdrawAddress
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgDisableServiceBinding, &weightMsgDisableServiceBinding, nil,
		func(_ *rand.Rand) {
			weightMsgDisableServiceBinding = simappparams.DefaultWeightMsgDisableServiceBinding
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgEnableServiceBinding, &weightMsgEnableServiceBinding, nil,
		func(_ *rand.Rand) {
			weightMsgEnableServiceBinding = simappparams.DefaultWeightMsgEnableServiceBinding
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRefundServiceDeposit, &weightMsgRefundServiceDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = simappparams.DefaultWeightMsgRefundServiceDeposit
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgDefineService,
			SimulateMsgDefineService(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgBindService,
			SimulateMsgBindService(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateServiceBinding,
			SimulateMsgUpdateServiceBinding(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgSetWithdrawAddress,
			SimulateMsgSetWithdrawAddress(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgDisableServiceBinding,
			SimulateMsgDisableServiceBinding(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgEnableServiceBinding,
			SimulateMsgEnableServiceBinding(ak, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRefundServiceDeposit,
			SimulateMsgRefundServiceDeposit(ak, k),
		),
	}
}

// SimulateMsgDefineService generates a MsgDefineService with random values.
func SimulateMsgDefineService(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)

		serviceName := simulation.RandStringOfLength(r, 20)
		serviceDescription := simulation.RandStringOfLength(r, 50)
		authorDescription := simulation.RandStringOfLength(r, 50)
		tags := []string{simulation.RandStringOfLength(r, 20), simulation.RandStringOfLength(r, 20)}
		schemas := `{"input":{"type":"object"},"output":{"type":"object"}}`

		account := ak.GetAccount(ctx, simAccount.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgDefineService(serviceName, serviceDescription, tags, simAccount.Address, authorDescription, schemas)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err := app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgBindService generates a MsgBindService with random values.
func SimulateMsgBindService(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)

		serviceName := simulation.RandStringOfLength(r, 20)
		deposit := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simulation.RandIntBetween(r, 1000000, 2000000)))))
		pricing := fmt.Sprintf(`{"price":"%d%s"`, simulation.RandIntBetween(r, 100, 1000), sdk.DefaultBondDenom)
		qos := uint64(simulation.RandIntBetween(r, 10, 100))

		account := ak.GetAccount(ctx, simAccount.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgBindService(serviceName, simAccount.Address, deposit, pricing, qos, simAccount.Address)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err := app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgUpdateServiceBinding generates a MsgUpdateServiceBinding with random values.
func SimulateMsgUpdateServiceBinding(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)

		serviceName := simulation.RandStringOfLength(r, 20)
		deposit := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simulation.RandIntBetween(r, 100, 1000)))))
		pricing := fmt.Sprintf(`{"price":"%d%s"`, simulation.RandIntBetween(r, 100, 1000), sdk.DefaultBondDenom)
		qos := uint64(simulation.RandIntBetween(r, 10, 100))

		account := ak.GetAccount(ctx, simAccount.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgUpdateServiceBinding(serviceName, simAccount.Address, deposit, pricing, qos, simAccount.Address)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err := app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgSetWithdrawAddress generates a MsgSetWithdrawAddress with random values.
func SimulateMsgSetWithdrawAddress(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)
		withdrawalAccount, _ := simulation.RandomAcc(r, accs)

		account := ak.GetAccount(ctx, simAccount.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgSetWithdrawAddress(simAccount.Address, withdrawalAccount.Address)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err := app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgDisableServiceBinding generates a MsgDisableServiceBinding with random values.
func SimulateMsgDisableServiceBinding(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)
		serviceName := simulation.RandStringOfLength(r, 20)

		account := ak.GetAccount(ctx, simAccount.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgDisableServiceBinding(serviceName, simAccount.Address, simAccount.Address)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err := app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgEnableServiceBinding generates a MsgEnableServiceBinding with random values.
func SimulateMsgEnableServiceBinding(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)

		serviceName := simulation.RandStringOfLength(r, 20)
		deposit := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simulation.RandIntBetween(r, 100, 1000)))))

		account := ak.GetAccount(ctx, simAccount.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgEnableServiceBinding(serviceName, simAccount.Address, deposit, simAccount.Address)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err := app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgRefundServiceDeposit generates a MsgRefundServiceDeposit with random values.
func SimulateMsgRefundServiceDeposit(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)
		serviceName := simulation.RandStringOfLength(r, 20)

		account := ak.GetAccount(ctx, simAccount.Address)
		fees, err := simulation.RandomFees(r, ctx, account.SpendableCoins(ctx.BlockTime()))
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgRefundServiceDeposit(serviceName, simAccount.Address, simAccount.Address)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err := app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}
