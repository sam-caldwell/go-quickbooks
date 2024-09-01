package quickbooks

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint pretty-prints the JSON form of the Account struct.
func (acct *Account) PrettyPrint() ([]byte, error) {
	prettyJson, err := json.MarshalIndent(acct, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to generate pretty JSON: %w", err)
	}
	return prettyJson, nil
}
