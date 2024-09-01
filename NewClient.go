package quickbooks

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	productionBaseUrl = "https://quickbooks.api.intuit.com/v3/company/%s/"
	developerBaseUrl  = "https://sandbox-quickbooks.api.intuit.com/v3/company/%s/"
)

// NewClient creates a new QuickBooks API client or return nil and an error
func NewClient(companyId, accessToken string, production bool) (*Client, error) {

	var baseUrl string
	if production {
		baseUrl = productionBaseUrl
	} else {
		baseUrl = developerBaseUrl
	}

	if strings.TrimSpace(accessToken) == "" {
		return nil, fmt.Errorf("token required")
	}
	if strings.TrimSpace(companyId) == "" {
		return nil, fmt.Errorf("company id required")
	}

	return &Client{
		BaseURL:     fmt.Sprintf(baseUrl, companyId),
		RealmID:     companyId,
		AccessToken: accessToken,
		HTTPClient:  &http.Client{Timeout: 10 * time.Second},
	}, nil
}
