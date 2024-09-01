package quickbooks

import (
	"encoding/json"
)

// Serialize converts the Account structure to a JSON-encoded []byte slice.
func (acct *Account) Serialize() ([]byte, error) {
	output, err := json.Marshal(acct)
	if err != nil {
		return nil, err
	}
	return output, nil
}
