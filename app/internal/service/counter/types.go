package counter

import "time"

type CounterInfo struct {
	Count           int32
	LastIncrementBy string
	LastIncrementAt time.Time
	NextFinalizeAt  time.Time
}

func (c *CounterInfo) LastIncrementAtTimestamp() string {
	if c.LastIncrementAt.IsZero() {
		return ""
	}
	return c.LastIncrementAt.String()
}

func (c *CounterInfo) NextFinalizeAtTimestamp() string {
	if c.NextFinalizeAt.IsZero() {
		return ""
	}
	return c.NextFinalizeAt.String()
}
