package score

import "errors"

var (
	ErrListScores  = errors.New("store list scores")
	ErrInsertScore = errors.New("store insert score")
)
