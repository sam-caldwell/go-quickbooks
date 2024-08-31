package quickbooks

// LineDetail contains details about a line item in a Journal Entry.
type LineDetail struct {
	PostingType string    `json:"PostingType"`
	AccountRef  Reference `json:"AccountRef"`
}
