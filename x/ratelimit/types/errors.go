package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrRateLimitExceeded = sdkerrors.Register(ModuleName, 2, "rate limit exceeded")
	ErrBadMessage        = sdkerrors.Register(ModuleName, 3, "bad message")
	ErrBadContract       = sdkerrors.Register(ModuleName, 4, "bad contract")
)
