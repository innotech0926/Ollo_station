package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// liquidity module sentinel errors
var (
	ErrInsufficientDepositAmount    = sdkerrors.Register(ModuleName, 2, "insufficient deposit amount")
	ErrPairAlreadyExists            = sdkerrors.Register(ModuleName, 3, "pair already exists")
	ErrWrongPoolCoinDenom           = sdkerrors.Register(ModuleName, 5, "wrong pool coin denom")
	ErrInvalidCoinDenom             = sdkerrors.Register(ModuleName, 6, "invalid coin denom")
	ErrNoLastPrice                  = sdkerrors.Register(ModuleName, 8, "cannot make a market order to a pair with no last price")
	ErrInsufficientOfferCoin        = sdkerrors.Register(ModuleName, 9, "insufficient offer coin")
	ErrPriceOutOfRange              = sdkerrors.Register(ModuleName, 10, "price out of range limit")
	ErrTooLongOrderLifespan         = sdkerrors.Register(ModuleName, 11, "order lifespan is too long")
	ErrDisabledPool                 = sdkerrors.Register(ModuleName, 12, "disabled pool")
	ErrWrongPair                    = sdkerrors.Register(ModuleName, 13, "wrong denom pair")
	ErrSameBatch                    = sdkerrors.Register(ModuleName, 14, "cannot cancel an order within the same batch")
	ErrAlreadyCanceled              = sdkerrors.Register(ModuleName, 15, "the order is already canceled")
	ErrDuplicatePairId              = sdkerrors.Register(ModuleName, 16, "duplicate pair id presents in the pair id list")
	ErrTooSmallOrder                = sdkerrors.Register(ModuleName, 17, "too small order")
	ErrTooLargePool                 = sdkerrors.Register(ModuleName, 18, "too large pool")
	ErrTooManyPools                 = sdkerrors.Register(ModuleName, 19, "too many pools in the pair")
	ErrPriceNotOnTicks              = sdkerrors.Register(ModuleName, 20, "price is not on ticks")
	ErrPoolNotExists                = sdkerrors.Register(ModuleName, 21, "pool not exists")
	ErrPoolTypeNotExists            = sdkerrors.Register(ModuleName, 22, "pool type not exists")
	ErrEqualDenom                   = sdkerrors.Register(ModuleName, 23, "reserve coin denomination are equal")
	ErrInvalidDenom                 = sdkerrors.Register(ModuleName, 24, "invalid denom")
	ErrNumOfReserveCoin             = sdkerrors.Register(ModuleName, 25, "invalid number of reserve coin")
	ErrNumOfPoolCoin                = sdkerrors.Register(ModuleName, 26, "invalid number of pool coin")
	ErrInsufficientPool             = sdkerrors.Register(ModuleName, 27, "insufficient pool")
	ErrInsufficientBalance          = sdkerrors.Register(ModuleName, 28, "insufficient coin balance")
	ErrLessThanMinInitDeposit       = sdkerrors.Register(ModuleName, 29, "deposit coin less than MinInitDepositAmount")
	ErrNotImplementedYet            = sdkerrors.Register(ModuleName, 30, "not implemented yet")
	ErrPoolAlreadyExists            = sdkerrors.Register(ModuleName, 31, "the pool already exists")
	ErrPoolBatchNotExists           = sdkerrors.Register(ModuleName, 32, "pool batch not exists")
	ErrOrderBookInvalidity          = sdkerrors.Register(ModuleName, 33, "orderbook is not validity")
	ErrBatchNotExecuted             = sdkerrors.Register(ModuleName, 34, "the liquidity pool batch is not executed")
	ErrInvalidPoolCreatorAddr       = sdkerrors.Register(ModuleName, 35, "invalid pool creator address")
	ErrInvalidDepositorAddr         = sdkerrors.Register(ModuleName, 36, "invalid pool depositor address")
	ErrInvalidWithdrawerAddr        = sdkerrors.Register(ModuleName, 37, "invalid pool withdrawer address")
	ErrInvalidSwapRequesterAddr     = sdkerrors.Register(ModuleName, 38, "invalid pool swap requester address")
	ErrBadPoolCoinAmount            = sdkerrors.Register(ModuleName, 39, "invalid pool coin amount")
	ErrBadDepositCoinsAmount        = sdkerrors.Register(ModuleName, 40, "invalid deposit coins amount")
	ErrBadOfferCoinAmount           = sdkerrors.Register(ModuleName, 41, "invalid offer coin amount")
	ErrBadOrderingReserveCoin       = sdkerrors.Register(ModuleName, 42, "reserve coin denoms not ordered alphabetical")
	ErrBadOrderPrice                = sdkerrors.Register(ModuleName, 43, "invalid order price")
	ErrNumOfReserveCoinDenoms       = sdkerrors.Register(ModuleName, 44, "invalid reserve coin denoms")
	ErrEmptyReserveAccountAddress   = sdkerrors.Register(ModuleName, 45, "empty reserve account address")
	ErrEmptyPoolCoinDenom           = sdkerrors.Register(ModuleName, 46, "empty pool coin denom")
	ErrBadOrderingReserveCoinDenoms = sdkerrors.Register(ModuleName, 47, "bad ordering reserve coin denoms")
	ErrBadReserveAccountAddress     = sdkerrors.Register(ModuleName, 48, "bad reserve account address")
	ErrBadPoolCoinDenom             = sdkerrors.Register(ModuleName, 49, "bad pool coin denom")
	ErrInsufficientPoolCreationFee  = sdkerrors.Register(ModuleName, 50, "insufficient balances for pool creation fee")
	ErrExceededMaxOrderable         = sdkerrors.Register(ModuleName, 51, "can not exceed max order ratio of reserve coins that can be ordered at a order")
	ErrBadBatchMsgIndex             = sdkerrors.Register(ModuleName, 52, "bad msg index of the batch")
	ErrSwapTypeNotExists            = sdkerrors.Register(ModuleName, 53, "swap type not exists")
	ErrLessThanMinOfferAmount       = sdkerrors.Register(ModuleName, 54, "offer amount should be over 100 micro")
	ErrBadOfferCoinFee              = sdkerrors.Register(ModuleName, 55, "bad offer coin fee")
	ErrNotMatchedReserveCoin        = sdkerrors.Register(ModuleName, 56, "does not match the reserve coin of the pool")
	ErrBadPoolTypeID                = sdkerrors.Register(ModuleName, 57, "invalid index of the pool type")
	ErrExceededReserveCoinLimit     = sdkerrors.Register(ModuleName, 58, "can not exceed reserve coin limit amount")
	ErrDepletedPool                 = sdkerrors.Register(ModuleName, 59, "the pool is depleted of reserve coin, reinitializing is required by deposit")
	ErrCircuitBreakerEnabled        = sdkerrors.Register(ModuleName, 60, "circuit breaker is triggered")
	ErrOverflowAmount               = sdkerrors.Register(ModuleName, 61, "invalid amount that can cause overflow")
)
