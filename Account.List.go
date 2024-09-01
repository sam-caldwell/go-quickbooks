package quickbooks

import "fmt"

// List lists accounts in a paginated manner.
func (acct *Account) List(client *Client, startPosition, maxResults int) (AccountSet, error) {
	endpoint := fmt.Sprintf("/account?startPosition=%d&maxResults=%d", startPosition, maxResults)
	var result AccountSet
	err := client.request("GET", endpoint, nil, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
