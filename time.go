package minecraft

import (
	"fmt"
	"strings"
	"time"
)

const minecraftTimeLayout = "2006-01-02 15:04:05 -0700"

var nilTime = (time.Time{}).UnixNano()

// Time specific to how minecraft reads/writes it (Java)
type Time struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}

	t.Time, err = time.Parse(minecraftTimeLayout, s)
	return
}

// MarshalJSON implements json.Marshaler
func (t *Time) MarshalJSON() ([]byte, error) {
	if t.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(minecraftTimeLayout))), nil
}

// IsZero returns if the value is zero
func (t *Time) IsZero() bool {
	return t.UnixNano() == nilTime
}
