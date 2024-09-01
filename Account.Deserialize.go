package quickbooks

import (
	"bytes"
	"encoding/json"
)

// Deserialize - a given []byte input as the Account struct state
func (acct *Account) Deserialize(in []byte) (err error) {
	return json.NewDecoder(
		bytes.NewBuffer(in)).
		Decode(acct)
}
