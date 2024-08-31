package quickbooks

// Create creates a new Journal Entry in QuickBooks.
func (je *Journal) Create(client *Client, journal Journal) (*Journal, error) {
	var result Journal
	err := client.request("POST", "/journalentry", journal, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
