package quickbooks

import "net/http"

// Client - represents the QuickBooks API client
type Client struct {
	BaseURL     string
	RealmID     string
	AccessToken string
	HTTPClient  *http.Client
}
