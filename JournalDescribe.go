package quickbooks

import "fmt"

// Describe retrieves the details of a Journal Entry by its ID.
func (je *Journal) Describe(client *Client, journalId string) (*Journal, error) {
	endpoint := fmt.Sprintf("/journalentry/%s", journalId)
	var result Journal
	err := client.request("GET", endpoint, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
