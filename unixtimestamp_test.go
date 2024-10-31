package gotime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnixTimestamp_MarshalJSON(t *testing.T) {
	now := NowTs()

	s := struct {
		T1 UnixTimestamp `json:"t1"`
		T2 UnixTimestamp
		T3 UnixTimestamp `json:"x"`
	}{
		ZeroTs().AddDays(356),
		ZeroTs().AddDays(10),
		now,
	}
	if dat, err := json.Marshal(&s); err != nil {
		t.Error(err)
	} else {
		require := fmt.Sprintf(`{"t1":30758400,"T2":864000,"x":%d}`, now)
		if string(dat) != require {
			t.Errorf("expect %s, got %s", require, string(dat))
		}
	}
}
