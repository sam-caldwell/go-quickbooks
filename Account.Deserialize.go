package quickbooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sam-caldwell/file"
	"os"
)

// Deserialize reads JSON data from a file and decodes it into the Account struct.
func (acct *Account) Deserialize(fileName string) error {
	if !file.Exists(fileName) {
		return fmt.Errorf("file does not exist: %s", fileName)
	}
	in, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.NewDecoder(bytes.NewBuffer(in)).Decode(acct)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}
