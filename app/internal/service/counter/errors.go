package counter

import (
	"errors"
	"time"
)

var (
	ErrDBConn = errors.New("db conn")

	ErrGetCounter                = errors.New("store get counter")
	ErrInsertIncrementRequest    = errors.New("store insert increment request")
	ErrUpdateCounterFinalizeTime = errors.New("store update counter finalize time")

	ErrEnqueueFinalizeIncrement = errors.New("enqueue finalize increment job")
)

type MultipleIncrementRequestError struct {
	User           string
	FinalizeWindow time.Duration
}

func (e *MultipleIncrementRequestError) Error() string {
	return "multiple increment request by user in finalize window"
}
