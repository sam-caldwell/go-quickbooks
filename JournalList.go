package quickbooks

import "fmt"

// List lists Journal Entries in a paginated manner.
func (je *Journal) List(client *Client, startPosition, maxResults int) ([]Journal, error) {
	endpoint := fmt.Sprintf("/journalentry?startPosition=%d&maxResults=%d", startPosition, maxResults)
	var result struct {
		JournalEntries []Journal `json:"Journal"`
	}
	err := client.request("GET", endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.JournalEntries, nil
}
