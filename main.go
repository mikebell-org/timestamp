package timestamp

import (
	"strconv"
	"time"
)

type Timestamp float64

const nano = 1000000000.0

func Now() Timestamp {
	return New(time.Now())
}

func New(t time.Time) Timestamp {
	return Timestamp(Timestamp(t.UnixNano()) / nano)
}

func (t Timestamp) Time() time.Time {
	sec := int64(t)
	nsec := int64((t - Timestamp(sec)) * nano)
	return time.Unix(sec, nsec)
}

func (t Timestamp) String() string {
	return t.Time().String()
}

func (t *Timestamp) UnmarshalJSON(b []byte) (err error) {
	var f float64
	if len(b) == 0 {
		return
	}
	if b[0] == '"' {
		b = b[1 : len(b)-1]
	}
	if f, err = strconv.ParseFloat(string(b), 10); err != nil {
		return err
	}
	*t = Timestamp(f)
	return nil
}
