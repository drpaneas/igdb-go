package igdbclient

import (
	"fmt"

	multierr "github.com/hashicorp/go-multierror"
	"github.com/parnurzeal/gorequest"
)

const (
	// BaseURLV4 is version 4
	BaseURLV4 = "https://api.igdb.com/v4"
)

// Client is the actual client
type Client struct {
	BaseURL    string
	accesToken string
	clientID   string
	cl         *gorequest.SuperAgent
}

// NewClient is a factory pattern
func NewClient(accesToken, clientID string) (*Client, error) {
	if accesToken == "" {
		return nil, fmt.Errorf("No API Access Token passed")
	}
	if clientID == "" {
		return nil, fmt.Errorf("No client ID is passed")
	}
	return &Client{
		BaseURL:    BaseURLV4,
		accesToken: accesToken,
		clientID:   clientID,
		cl:         gorequest.New(),
	}, nil
}

// Response from searching for Trino game
type Response []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SearchGame sends a request for game search
func (c *Client) SearchGame(name string) (*Response, error) {
	SearchGameURL := fmt.Sprintf("%s/games", c.BaseURL)
	resp := &Response{}
	httpRes, _, errs := c.cl.Get(SearchGameURL).
		Set("Content-Type", "applications/json").
		Set("Accept", "application/json").
		Set("Client-ID", c.clientID).
		Set("Authorization", fmt.Sprintf("Bearer %s", c.accesToken)).
		Param("search", name).
		Param("fields", "name").
		EndStruct(resp)
	// dst := &bytes.Buffer{}
	// if err := json.Indent(dst, body, "", "  "); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(dst.String())
	if len(errs) > 0 {
		return nil, &multierr.Error{Errors: errs}
	}
	if httpRes.StatusCode != 200 {
		return nil, fmt.Errorf("http Status Code %d returned", httpRes.StatusCode)
	}
	return resp, nil
}
