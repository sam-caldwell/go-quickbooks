package quickbooks

// Create creates a new account in QuickBooks.
func (acct *Account) Create(client *Client, account Account) (*Account, error) {
	var result Account
	err := client.request("POST", "/account", account, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
