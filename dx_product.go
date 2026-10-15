package trivec

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// UnmarshalJSON decodes the VAT percentages of a product leniently.
//
// Trivec changed the contents of `Vat` on some properties (CSI-1650): it used
// to be an object (see DxVATCode), but is a plain number on others. `Vat2` is a
// plain number. Both are null for products without VAT. A shape that isn't
// known gives a nil percentage instead of failing the whole product list.
func (p *DxProduct) UnmarshalJSON(b []byte) error {
	type Alias DxProduct
	aux := &struct {
		*Alias
		Vat  json.RawMessage `json:"Vat"`
		Vat2 json.RawMessage `json:"Vat2"`
	}{Alias: (*Alias)(p)}

	err := json.Unmarshal(b, aux)
	if err != nil {
		return err
	}

	p.Vat = parseVatPercentage(aux.Vat)
	p.Vat2 = parseVatPercentage(aux.Vat2)
	return nil
}

// parseVatPercentage reads a VAT percentage from a number, a numeric string, or
// an object with a `Percentage`. It returns nil if there is no percentage.
func parseVatPercentage(raw json.RawMessage) *float64 {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 || bytes.Equal(raw, []byte("null")) {
		return nil
	}

	var number float64
	if err := json.Unmarshal(raw, &number); err == nil {
		return &number
	}

	var str string
	if err := json.Unmarshal(raw, &str); err == nil {
		number, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil
		}
		return &number
	}

	var object struct {
		Percentage *float64 `json:"Percentage"`
	}
	if err := json.Unmarshal(raw, &object); err == nil {
		return object.Percentage
	}

	return nil
}
