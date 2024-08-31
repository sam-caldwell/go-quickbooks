package quickbooks

// Line represents a line item in a Journal Entry.
type Line struct {
	ID                string     `json:"Id,omitempty"`
	Amount            float64    `json:"Amount"`
	Description       string     `json:"Description,omitempty"`
	JournalLineDetail LineDetail `json:"JournalLineDetail"`
}
