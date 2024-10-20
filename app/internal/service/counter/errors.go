package counter

import (
	"errors"
	"time"
)

var (
	ErrGetCounter                = errors.New("store get counter")
	ErrIncrementCounter          = errors.New("store increment counter")
	ErrUpdateCounterFinalizeTime = errors.New("store update counter finalize time")
	ErrResetCounter              = errors.New("store reset counter")

	ErrListIncrementRequests     = errors.New("store list increment requests")
	ErrInsertIncrementRequest    = errors.New("store insert increment request")
	ErrTruncateIncrementRequests = errors.New("store truncate increment requests")

	ErrEnqueueFinalizeIncrement = errors.New("enqueue finalize increment job")
)

type MultipleIncrementRequestError struct {
	User           string
	FinalizeWindow time.Duration
}

func (e *MultipleIncrementRequestError) Error() string {
	return "multiple increment request by user in finalize window"
}
