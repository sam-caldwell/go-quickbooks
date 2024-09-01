package quickbooks

import "fmt"

// Describe retrieves the details of an account matching the given accountId.
func (acct *Account) Describe(client *Client, accountId string) (*Account, error) {
	endpoint := fmt.Sprintf("/account/%s", accountId)
	err := client.request("GET", endpoint, nil, acct)
	if err != nil {
		return nil, err
	}
	return acct, nil
}
