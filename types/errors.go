package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// service module sentinel errors
var (
	ErrInvalidServiceName       = sdkerrors.Register(ModuleName, 1, "invalid service name, must contain alphanumeric characters, _ and - only，length greater than 0 and less than or equal to 70")
	ErrInvalidDescription       = sdkerrors.Register(ModuleName, 2, "invalid description")
	ErrInvalidTags              = sdkerrors.Register(ModuleName, 3, "invalid tags")
	ErrInvalidSchemas           = sdkerrors.Register(ModuleName, 4, "invalid schemas")
	ErrUnknownServiceDefinition = sdkerrors.Register(ModuleName, 5, "unknown service definition")
	ErrServiceDefinitionExists  = sdkerrors.Register(ModuleName, 6, "service definition already exists")

	ErrInvalidDeposit            = sdkerrors.Register(ModuleName, 7, "invalid deposit")
	ErrInvalidPricing            = sdkerrors.Register(ModuleName, 8, "invalid pricing")
	ErrInvalidQoS                = sdkerrors.Register(ModuleName, 9, "invalid qos")
	ErrServiceBindingExists      = sdkerrors.Register(ModuleName, 10, "service binding already exists")
	ErrUnknownServiceBinding     = sdkerrors.Register(ModuleName, 11, "unknown service binding")
	ErrServiceBindingUnavailable = sdkerrors.Register(ModuleName, 12, "service binding unavailable")
	ErrServiceBindingAvailable   = sdkerrors.Register(ModuleName, 13, "service binding available")
	ErrIncorrectRefundTime       = sdkerrors.Register(ModuleName, 14, "incorrect refund time")

	ErrInvalidServiceFee         = sdkerrors.Register(ModuleName, 15, "invalid service fee")
	ErrInvalidProviders          = sdkerrors.Register(ModuleName, 16, "invalid providers")
	ErrInvalidTimeout            = sdkerrors.Register(ModuleName, 17, "invalid timeout")
	ErrInvalidRepeatedFreq       = sdkerrors.Register(ModuleName, 18, "invalid repeated frequency")
	ErrInvalidRepeatedTotal      = sdkerrors.Register(ModuleName, 19, "invalid repeated total count")
	ErrInvalidResponseThreshold  = sdkerrors.Register(ModuleName, 20, "invalid response threshold")
	ErrInvalidResponse           = sdkerrors.Register(ModuleName, 21, "invalid response")
	ErrInvalidRequestID          = sdkerrors.Register(ModuleName, 22, "invalid request ID")
	ErrUnknownRequest            = sdkerrors.Register(ModuleName, 23, "unknown request")
	ErrUnknownResponse           = sdkerrors.Register(ModuleName, 24, "unknown response")
	ErrUnknownRequestContext     = sdkerrors.Register(ModuleName, 25, "unknown request context")
	ErrInvalidRequestContextID   = sdkerrors.Register(ModuleName, 26, "invalid request context ID")
	ErrRequestContextNonRepeated = sdkerrors.Register(ModuleName, 27, "request context non repeated")
	ErrRequestContextNotRunning  = sdkerrors.Register(ModuleName, 28, "request context not running")
	ErrRequestContextNotPaused   = sdkerrors.Register(ModuleName, 29, "request context not paused")
	ErrRequestContextCompleted   = sdkerrors.Register(ModuleName, 30, "request context completed")
	ErrCallbackRegistered        = sdkerrors.Register(ModuleName, 31, "callback registered")
	ErrCallbackNotRegistered     = sdkerrors.Register(ModuleName, 32, "callback not registered")
	ErrInvalidEarnedFees         = sdkerrors.Register(ModuleName, 33, "no earned fees or being zero")

	ErrInvalidRequestInput   = sdkerrors.Register(ModuleName, 34, "invalid request input")
	ErrInvalidResponseOutput = sdkerrors.Register(ModuleName, 35, "invalid response output")
	ErrInvalidResponseResult = sdkerrors.Register(ModuleName, 36, "invalid response result")

	ErrInvalidSchemaName = sdkerrors.Register(ModuleName, 37, "invalid service schema name")
	ErrNotAuthorized     = sdkerrors.Register(ModuleName, 38, "not authorized")
)
