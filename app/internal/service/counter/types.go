package counter

import "time"

type Info struct {
	Count           int32
	LastIncrementBy string
	LastIncrementAt time.Time
	NextFinalizeAt  time.Time
}

func (i *Info) LastIncrementAtTimestamp() string {
	if i.LastIncrementAt.IsZero() {
		return ""
	}
	return i.LastIncrementAt.String()
}

func (i *Info) NextFinalizeAtTimestamp() string {
	if i.NextFinalizeAt.IsZero() {
		return ""
	}
	return i.NextFinalizeAt.String()
}
