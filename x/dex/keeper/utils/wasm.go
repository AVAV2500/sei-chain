package utils

import (
	"encoding/json"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/utils"
	"github.com/sei-protocol/sei-chain/utils/metrics"
	"github.com/sei-protocol/sei-chain/x/dex/keeper"
	dextypeswasm "github.com/sei-protocol/sei-chain/x/dex/types/wasm"
)

const ErrWasmModuleInstCPUFeatureLiteral = "Error instantiating module: CpuFeature"

func getMsgType(msg interface{}) string {
	switch msg.(type) {
	case dextypeswasm.SudoNewBlockMsg:
		return "new_block"
	case dextypeswasm.SudoFinalizeBlockMsg:
		return "finalize_block"
	case *dextypeswasm.SudoFinalizeBlockMsg:
		return "finalize_block"
	case dextypeswasm.SudoSettlementMsg:
		return "settlement"
	case dextypeswasm.SudoOrderPlacementMsg:
		return "bulk_order_placements"
	case dextypeswasm.SudoOrderCancellationMsg:
		return "bulk_order_cancellations"
	case dextypeswasm.SudoLiquidationMsg:
		return "liquidation"
	default:
		return "unknown"
	}
}

func sudo(sdkCtx sdk.Context, k *keeper.Keeper, contractAddress []byte, wasmMsg []byte, msgType string) ([]byte, uint64, error) {
	// Measure the time it takes to execute the contract in WASM
	defer metrics.MeasureSudoExecutionDuration(time.Now(), msgType)
	gasConsumedBefore := sdkCtx.GasMeter().GasConsumed()
	data, err := k.WasmKeeper.Sudo(
		sdkCtx, contractAddress, wasmMsg,
	)
	gasConsumedAfter := sdkCtx.GasMeter().GasConsumed()
	if hasErrInstantiatingWasmModuleDueToCPUFeature(err) {
		panic(utils.DecorateHardFailError(err))
	}
	return data, gasConsumedAfter - gasConsumedBefore, err
}

func hasErrInstantiatingWasmModuleDueToCPUFeature(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), ErrWasmModuleInstCPUFeatureLiteral)
}

func CallContractSudo(sdkCtx sdk.Context, k *keeper.Keeper, contractAddr string, msg interface{}) ([]byte, error) {
	contractAddress, err := sdk.AccAddressFromBech32(contractAddr)
	if err != nil {
		sdkCtx.Logger().Error(err.Error())
		return []byte{}, err
	}
	wasmMsg, err := json.Marshal(msg)
	if err != nil {
		sdkCtx.Logger().Error(err.Error())
		return []byte{}, err
	}
	msgType := getMsgType(msg)
	data, gasUsed, err := sudo(sdkCtx, k, contractAddress, wasmMsg, msgType)
	if err != nil {
		metrics.IncrementSudoFailCount(msgType)
		sdkCtx.Logger().Error(err.Error())
		return []byte{}, err
	}
	if err := k.ChargeRentForGas(sdkCtx, contractAddr, gasUsed); err != nil {
		metrics.IncrementSudoFailCount(msgType)
		sdkCtx.Logger().Error(err.Error())
		return []byte{}, err
	}
	return data, nil
}
