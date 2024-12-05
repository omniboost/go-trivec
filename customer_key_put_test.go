package trivec_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerKeyPut(t *testing.T) {
	req := client.NewCustomerKeyPutRequest()
	req.PathParams().AccountKey = "3526138085244930"
	req.RequestBody().ClientNr =  "123"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}



