package igdb

import (
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.igdb.com/"
	userAgent      = "go-igdb"
)

// A Client manages communication with the IGBD API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// BASE URL for API requests. Defaults to the public IGDB API.
	// It should always be specified with a trailing slash.
	BaseURL *url.URL

	// UserAgent used when communicating with IGDB API.
	UserAgent string
}
