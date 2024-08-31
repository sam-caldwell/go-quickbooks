package quickbooks

import "fmt"

// Delete deletes an account by its accountId.
func (acct *Account) Delete(client *Client, accountId string) error {
	endpoint := fmt.Sprintf("/account/%s?operation=delete", accountId)
	if err := client.request("POST", endpoint, nil, nil); err != nil { // Use "POST" for delete operation as required by QuickBooks
		return err
	}
	return nil
}
