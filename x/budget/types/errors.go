package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Sentinel errors for the budget module.
var (
	ErrInvalidBudgetName      = sdkerrors.Register(ModuleName, 2, "budget name only allows letters, digits, and dash(-) without spaces and the maximum length is 50")
	ErrInvalidStartEndTime    = sdkerrors.Register(ModuleName, 3, "budget end time should not be earlier than budget start time")
	ErrInvalidBudgetRate      = sdkerrors.Register(ModuleName, 4, "invalid budget rate")
	ErrInvalidTotalBudgetRate = sdkerrors.Register(ModuleName, 5, "invalid total rate of the budgets with the same budget source address")
	ErrDuplicateBudgetName    = sdkerrors.Register(ModuleName, 6, "duplicate budget name")
)
