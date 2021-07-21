package trivec_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	trivec "github.com/omniboost/go-trivec"
)

var (
	client *trivec.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	subscriptionKey := os.Getenv("SUBSCRIPTION_KEY")
	serviceKey := os.Getenv("SERVICE_KEY")
	serviceID := os.Getenv("SERVICE_ID")
	appID := os.Getenv("APP_ID")
	environment := os.Getenv("ENVIRONMENT")
	debug := os.Getenv("DEBUG")

	client = trivec.NewClient(nil, subscriptionKey, serviceKey, serviceID, appID)
	if debug != "" {
		client.SetDebug(true)
	}
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	if environment != "" {
		client.SetEnvironment(environment)
	}
	m.Run()
}
