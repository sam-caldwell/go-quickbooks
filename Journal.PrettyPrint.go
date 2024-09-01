package quickbooks

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint - pretty-prints the JSON form of the JournalEntry struct.
func (je *Journal) PrettyPrint() ([]byte, error) {
	prettyJson, err := json.MarshalIndent(je, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to generate pretty JSON: %w", err)
	}
	return prettyJson, nil
}
