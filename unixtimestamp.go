package gotime

import (
	"strconv"
	"time"
)

type UnixTimestamp uint32

func From(t time.Time) UnixTimestamp {
	return UnixTimestamp(t.Unix())
}

func NowTs() UnixTimestamp {
	return UnixTimestamp(time.Now().Unix())
}

func (t UnixTimestamp) AddMinutes(minutes int) UnixTimestamp {
	return UnixTimestamp(uint32(t) + uint32(minutes*60))
}

func (t UnixTimestamp) AddDays(days int) UnixTimestamp {
	return UnixTimestamp(uint32(t) + uint32(days*3600*24))
}

func (t UnixTimestamp) IsSet() bool {
	return t > 0
}

func (t UnixTimestamp) IsFuture() bool {
	return t > NowTs()
}

func (t UnixTimestamp) IsPast() bool {
	return t <= NowTs()
}

func (t UnixTimestamp) NilIfZero() interface{} {
	if ZeroTs() == t {
		return nil
	}
	return t.Uint32()
}

func (t UnixTimestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}
func (t UnixTimestamp) String() string {
	return strconv.Itoa(int(t.Uint32()))
}
func (t UnixTimestamp) Uint32() uint32 {
	return uint32(t)
}

func ZeroTs() UnixTimestamp {
	return UnixTimestamp(0)
}

func MinTs(a, b UnixTimestamp) UnixTimestamp {
	if a < b {
		return a
	}

	return b
}

func MaxTs(a, b UnixTimestamp) UnixTimestamp {
	if a > b {
		return a
	}

	return b
}
