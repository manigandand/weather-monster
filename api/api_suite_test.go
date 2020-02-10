package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"weather-monster/api/v1"
	. "weather-monster/api"
	"weather-monster/config"
	"weather-monster/pkg/trace"
	"weather-monster/schema"

	"github.com/go-chi/chi"

	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}

var _ = BeforeSuite(func() {
	// connect to database
	MainSetup()
})

var _ = AfterSuite(func() {
	MainTearDown()
})

// Test setup ------------------------------------------------------------

// MainSetup common test setup, db, redis, env, and migrations
func MainSetup() {
	config.Initialize()
	Setup()

	router := chi.NewRouter()
	router.Route("/", v1.Routes)
	if tServer == nil {
		tServer = httptest.NewServer(router)
	}
	if tClient == nil {
		tClient = NewClient(tServer.URL)
	}

	InitService("test", "1")
	trace.Setup(config.Env)
}

// Setup overide the db credentials with test credentials
func Setup() {
	config.DBDriver = config.TestDBDriver
	config.DBDataSource = config.TestDBDataSource

	
}

// MainTearDown ...
func MainTearDown() {
	if tServer != nil {
		tServer.Close()
	}
	TearDown()
}

// TearDown ...
func TearDown() {
	db := Store.DB.Unscoped()
	db.Delete(schema.City{})
	db.Delete(schema.Temperature{})
	db.Delete(schema.Webhook{})
}

// Test setup ------------------------------------------------------------

var (
	tServer *httptest.Server
	tClient *Client
)

// Client ...
type Client struct {
	URL        string
	HTTPClient *http.Client
}

// NewClient ...
func NewClient(URL string) *Client {
	return &Client{URL, &http.Client{}}
}

// DoGet makes http.get req
func (c *Client) DoGet(url string) (*http.Response, error) {
	return c.Do("GET", c.URL+url, nil)
}

// Do makes http request
func (c *Client) Do(method, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
