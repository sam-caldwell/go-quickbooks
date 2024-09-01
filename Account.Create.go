package quickbooks

// Create creates a new account in QuickBooks.
func (acct *Account) Create(client *Client) (*Account, error) {
	var result Account
	err := client.request("POST", "/account", acct, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
