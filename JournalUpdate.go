package quickbooks

import "fmt"

// Update updates an existing Journal Entry in QuickBooks.
func (je *Journal) Update(client *Client, journal Journal) (*Journal, error) {
	if journal.ID == "" {
		return nil, fmt.Errorf("journal entry ID is required for update")
	}
	endpoint := fmt.Sprintf("/journalentry/%s", journal.ID)
	var result Journal
	err := client.request("POST", endpoint, journal, &result) // POST with payload for update
	if err != nil {
		return nil, err
	}
	return &result, nil
}
