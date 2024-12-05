package trivec_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestOpenSessionPost(t *testing.T) {
	req := client.NewOpenSessionPostRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}



