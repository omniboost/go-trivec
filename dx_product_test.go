package trivec_test

import (
	"encoding/json"
	"testing"

	trivec "github.com/omniboost/go-trivec"
)

func TestDxProductVat(t *testing.T) {
	f := func(v float64) *float64 { return &v }

	tests := []struct {
		name string
		json string
		vat  *float64
		vat2 *float64
	}{
		{"numbers", `{"Key":"1","Vat":12.0,"Vat2":6.0}`, f(12), f(6)},
		{"zero is not nothing", `{"Key":"1","Vat":0,"Vat2":0}`, f(0), f(0)},
		{"nulls", `{"Key":"1","Vat":null,"Vat2":null}`, nil, nil},
		{"missing", `{"Key":"1"}`, nil, nil},
		{"object", `{"Key":"1","Vat":{"$type":"x","Code":2,"Default2":false,"Key":"k","Name":"Food","Percentage":12.0},"Vat2":6.0}`, f(12), f(6)},
		{"object without percentage", `{"Key":"1","Vat":{"Code":2},"Vat2":6.0}`, nil, f(6)},
		{"string", `{"Key":"1","Vat":"25","Vat2":"12.5"}`, f(25), f(12.5)},
		{"unknown shapes don't fail", `{"Key":"1","Vat":[1,2],"Vat2":true}`, nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p trivec.DxProduct
			if err := json.Unmarshal([]byte(tt.json), &p); err != nil {
				t.Fatal(err)
			}
			if p.Key != "1" {
				t.Errorf("other fields are not decoded: key = %q", p.Key)
			}
			check := func(name string, got, want *float64) {
				switch {
				case got == nil && want == nil:
				case got == nil || want == nil:
					t.Errorf("%s: got %v, want %v", name, got, want)
				case *got != *want:
					t.Errorf("%s: got %v, want %v", name, *got, *want)
				}
			}
			check("Vat", p.Vat, tt.vat)
			check("Vat2", p.Vat2, tt.vat2)
		})
	}
}

func TestDxProductsVat(t *testing.T) {
	// a list mixing the shapes must decode as a whole
	var ps trivec.DxProducts
	err := json.Unmarshal([]byte(`[{"Key":"a","Vat":12,"Vat2":6},{"Key":"b","Vat":{"Percentage":25},"Vat2":12},{"Key":"c","Vat":null,"Vat2":null}]`), &ps)
	if err != nil {
		t.Fatal(err)
	}
	if len(ps) != 3 || *ps[0].Vat != 12 || *ps[1].Vat != 25 || ps[2].Vat != nil {
		t.Errorf("unexpected products: %+v", ps)
	}
}
