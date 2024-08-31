package quickbooks

import "fmt"

// Describe retrieves the details of an account matching the given accountId.
func (acct *Account) Describe(client *Client, accountId string) (*Account, error) {
	endpoint := fmt.Sprintf("/account/%s", accountId)
	var result Account
	err := client.request("GET", endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
