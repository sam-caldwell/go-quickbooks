package quickbooks

import (
	"encoding/json"
)

// Serialize - serialize the Account structure to a []byte slice
func (acct *Account) Serialize() (output []byte, err error) {
	if output, err = json.Marshal(*acct); err != nil {
		return nil, err
	}
	return output, nil
}
