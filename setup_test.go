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
	baseURLExportServiceString := os.Getenv("BASE_URL_EXPORT_SERVICE")
	baseURLLiteAPIString := os.Getenv("BASE_URL_LITE_API")
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
	if baseURLExportServiceString != "" {
		baseURL, err := url.Parse(baseURLExportServiceString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURLExportService(*baseURL)
	}
	if baseURLLiteAPIString != "" {
		baseURL, err := url.Parse(baseURLLiteAPIString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURLLiteAPI(*baseURL)
	}
	if environment != "" {
		client.SetEnvironment(environment)
	}
	client.SetDisallowUnknownFields(true)
	m.Run()

	os.Exit(0)
}
