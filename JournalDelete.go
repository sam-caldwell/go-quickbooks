package quickbooks

import "fmt"

// Delete deletes a Journal Entry by its ID.
func (je *Journal) Delete(client *Client, journalId string) error {
	endpoint := fmt.Sprintf("/journalentry/%s?operation=delete", journalId)
	if err := client.request("POST", endpoint, nil, nil); err != nil { // Use POST for delete operation as per QBO API
		return err
	}
	return nil
}
