package quickbooks

import "fmt"

// List lists accounts in a paginated manner.
func (acct *Account) List(client *Client, startPosition, maxResults int) ([]Account, error) {
	endpoint := fmt.Sprintf("/account?startPosition=%d&maxResults=%d", startPosition, maxResults)
	var result struct {
		Accounts []Account `json:"Account"`
	}
	err := client.request("GET", endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Accounts, nil
}
