package entity

import "errors"

// ErrTimeoutExceeded timeout
var ErrTimeoutExceeded = errors.New("timeout exceeded")

// ErrMissingSourceOrTarget when missing source and/or target
var ErrMissingSourceOrTarget = errors.New("source and/or target missing")

// ErrInvalidPrice price cannot be less than equal to 0
var ErrInvalidPrice = errors.New("price cannot be less than equal to 0")

// ErrEmptyCode when the code is an empry string
var ErrEmptyCode = errors.New("empty Code")

// ErrLenCode when the code is > OR < to 3
var ErrLenCode = errors.New("the Code cannot be greater than 3 or less than 3")

// ErrInvalidCaseCode when the code is does not uppercase string
var ErrInvalidCaseCode = errors.New("the Code needs to be uppercase")
