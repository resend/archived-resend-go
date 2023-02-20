package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

var (
	_ json.Marshaler   = &Date{}
	_ json.Unmarshaler = &Date{}
	_ fmt.Stringer     = &Date{}
)

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format("2006-01-02"))), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var err error

	str := string(data)
	str = strings.Trim(str, `"`)

	d.Time, err = time.Parse("2006-01-02", str)
	return err
}

func (d Date) String() string {
	return d.Time.Format("2006-01-02")
}
