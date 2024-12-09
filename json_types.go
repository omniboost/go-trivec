package trivec

import (
	"encoding/json"
	"math"
	"strconv"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("20060102"))
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		var i int
		err = json.Unmarshal(text, &i)
		if err != nil {
			var f float64
			err = json.Unmarshal(text, &f)
			if err != nil {
				return err
			}

			i = int(math.Round(f))
			value = strconv.Itoa(i)
		}

		value = strconv.Itoa(i)
	}

	if value == "" {
		return nil
	}

	if value == "10101" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try datetime without time zone
	d.Time, err = time.Parse("2006-01-02T15:04:05", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("20060102", value)
	return
}

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}
