package webhooks_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"weather-monster/api"
	v1 "weather-monster/api/v1"
	"weather-monster/config"
	"weather-monster/pkg/trace"
	"weather-monster/schema"

	"testing"

	"github.com/go-chi/chi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWebhooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Webhooks Suite")
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
	api.InitService("test", "1")

	router := chi.NewRouter()
	router.Route("/", v1.Routes)
	if tServer == nil {
		tServer = httptest.NewServer(router)
	}
	if tClient == nil {
		tClient = NewClient(tServer.URL)
	}

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
	db := api.Store.DB.Unscoped()
	db.Delete(schema.City{})
	db.Delete(schema.Temperature{})
	db.Delete(schema.Webhook{})

	db.Exec("ALTER SEQUENCE cities_id_seq RESTART WITH 1;")
	db.Exec("ALTER SEQUENCE temperatures_id_seq RESTART WITH 1;")
	db.Exec("ALTER SEQUENCE webhooks_id_seq RESTART WITH 1;")
}

// Test setup ------------------------------------------------------------

var (
	tServer *httptest.Server
	tClient *Client
)

// Client ...
type Client struct {
	URL        string
	Version    string
	HTTPClient *http.Client
}

// NewClient ...
func NewClient(URL string) *Client {
	return &Client{URL, "/v1", &http.Client{}}
}

// DoGet makes http.get req
func (c *Client) DoGet(url string) (*http.Response, error) {
	return c.Do(http.MethodGet, c.URL+c.Version+url, nil)
}

// DoPost makes http.Post req
func (c *Client) DoPost(url string, body interface{}) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return c.Do(http.MethodPost, c.URL+c.Version+url, bytes.NewReader(b))
}

// DoDelete makes http.Delete req
func (c *Client) DoDelete(url string) (*http.Response, error) {
	return c.Do(http.MethodDelete, c.URL+c.Version+url, nil)
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

func (c *Client) PostCities(req *schema.CityReq) (*http.Response, error) {
	return c.DoPost("/cities", req)
}

func (c *Client) GetWebhook() (*http.Response, error) {
	return c.DoGet("/webhooks")
}

func (c *Client) PostWebhooks(req *schema.Webhook) (*http.Response, error) {
	return c.DoPost("/webhooks", req)
}

func (c *Client) GetWebhookByID(id int) (*http.Response, error) {
	return c.DoGet(fmt.Sprintf("/webhooks/%d", id))
}

func (c *Client) DeleteWebhook(id int) (*http.Response, error) {
	return c.DoDelete(fmt.Sprintf("/webhooks/%d", id))
}
