package trivec_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerPost(t *testing.T) {
	req := client.NewCustomerPostRequest()
	req.PathParams().ClientNr = "3526138085244930"
	req.RequestBody().ClientNr =  "123"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}


