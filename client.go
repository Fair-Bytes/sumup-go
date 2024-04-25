// Package sumup provides a client for using the SumUp API.
// Usage:
//
//	import "github.com/sumup/sumup-go"
//
// Construct a new SumUp client, then use the various services on the client to
// access different parts of the SumUp API. For example:
//
//	client := sumup.NewClient().WithAuth(os.Getenv("SUMUP_KEY"))
//
//	// get the account the client is currently authorized for
//	account, err := client.Merchant.Get(context.Background(), sumup.GetAccountParams{})
//
// The client is structured around individual services that correspond to the tags
// in SumUp documentation https://developer.sumup.com/docs/api/sum-up-rest-api/.
package sumup

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//go:embed .version
var version string

const (
	// APIUrl is the URL of our API. Currently, SumUp doesn't provide any
	// other environment or APIs thus APIUrl is used as the default URL
	// for the client.
	APIUrl = "https://api.sumup.com"
)

type service struct {
	client *Client
}

// Client manages communication with the SumUp APIs.
type Client struct {
	// service is the shared service struct re-used for all services.
	svc service

	// client is the HTTP client used to communicate with the API.
	client *http.Client
	// url is the url of the API the requests will be sent to.
	url string
	// userAgent is the user-agent header that will be sent with
	// every request.
	userAgent string
	// key is the API key or access token used for authorization.
	key string

	ApiKeys       *ApiKeysService
	Authorization *AuthorizationService
	Checkouts     *CheckoutsService
	Customers     *CustomersService
	Merchant      *MerchantService
	Payouts       *PayoutsService
	Receipts      *ReceiptsService
	Subaccounts   *SubaccountsService
	Transactions  *TransactionsService
}

// NewClient creates new SumUp API [Client].
// To use APIs that require authentication use [Client.WithAuth].
func NewClient() *Client {
	c := &Client{
		client:    http.DefaultClient,
		userAgent: fmt.Sprintf("sumup-go/%s", version),
		url:       APIUrl,
	}
	c.populate()
	return c
}

// WithAuth returns a copy of the [Client] configured with the provided Authorization key.
func (c *Client) WithAuth(key string) *Client {
	clone := Client{
		client:    c.client,
		url:       APIUrl,
		userAgent: c.userAgent,
		key:       key,
	}
	clone.populate()
	return &clone
}

// WithClient returns a copy of the [Client] configured with the provided http client.
func (c *Client) WithHTTPClient(client *http.Client) *Client {
	clone := Client{
		client:    client,
		url:       APIUrl,
		userAgent: c.userAgent,
		key:       c.key,
	}
	clone.populate()
	return &clone
}

func (c *Client) populate() {
	c.svc.client = c
	c.ApiKeys = (*ApiKeysService)(&c.svc)
	c.Authorization = (*AuthorizationService)(&c.svc)
	c.Checkouts = (*CheckoutsService)(&c.svc)
	c.Customers = (*CustomersService)(&c.svc)
	c.Merchant = (*MerchantService)(&c.svc)
	c.Payouts = (*PayoutsService)(&c.svc)
	c.Receipts = (*ReceiptsService)(&c.svc)
	c.Subaccounts = (*SubaccountsService)(&c.svc)
	c.Transactions = (*TransactionsService)(&c.svc)
}

func (c *Client) NewRequest(
	ctx context.Context,
	method, path string,
	body io.Reader,
) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.url+path,
		body,
	)
	if err != nil {
		return nil, fmt.Errorf("build request: %s", err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+c.key)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("SumUp-Version", version)
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
