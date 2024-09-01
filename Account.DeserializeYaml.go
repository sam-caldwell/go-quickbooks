package quickbooks

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// DeserializeYAML deserializes the YAML file into the Account struct
func (acct *Account) DeserializeYAML(fileName string) error {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = yaml.Unmarshal(fileContent, acct)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return nil
}
