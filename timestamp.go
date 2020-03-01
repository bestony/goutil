package goutil

import "time"

type Timestamp int64

func (t Timestamp) UTCTime() time.Time {
	return t.LocalTime().UTC()
}

func (t Timestamp) LocalTime() time.Time {
	return time.Unix(int64(t), 0)
}
