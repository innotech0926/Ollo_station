package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	ibcante "github.com/cosmos/ibc-go/v6/modules/core/ante"
	"github.com/cosmos/ibc-go/v6/modules/core/keeper"
	wasmkeeper "github.com/ollo-station/ollo/x/wasm/keeper"
	wasmtypes "github.com/ollo-station/ollo/x/wasm/types"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC keeper.
type HandlerOptions struct {
	ante.HandlerOptions
	// IBCKeeper         app.IBCKeeper;
	// WasmConfig        &wasmConfig;
	// TXCounterStoreKey keys[wasm.StoreKey];

	// ante.HandlerOptions{
	// 	AccountKeeper:   app.AccountKeeper,
	// 	BankKeeper:      app.BankKeeper,
	// 	FeegrantKeeper:  app.FeeGrantKeeper,
	// 	SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
	// 	SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
	// },
	//
	IBCKeeper         *keeper.Keeper
	WasmConfig        *wasmtypes.WasmConfig
	TXCounterStoreKey storetypes.StoreKey
}

// NewAnteHandler creates a new ante handler
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrLogic,
			"sign mode handler is required for AnteHandler",
		)
	}
	if options.WasmConfig == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "wasm config is required for ante builder")
	}
	if options.TXCounterStoreKey == nil {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrLogic,
			"tx counter key is required for ante builder",
		)
	}

	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}
	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first

		wasmkeeper.NewLimitSimulationGasDecorator(
			options.WasmConfig.SimulationGasLimit,
		), // after setup context to enforce limits early
		wasmkeeper.NewCountTXDecorator(options.TXCounterStoreKey),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(
			options.AccountKeeper,
			options.BankKeeper,
			options.FeegrantKeeper,
			options.TxFeeChecker,
		),
		ante.NewSetPubKeyDecorator(
			options.AccountKeeper,
		), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
