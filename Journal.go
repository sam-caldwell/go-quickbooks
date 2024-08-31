package quickbooks

// Journal represents a QuickBooks Journal Entry.
type Journal struct {
	ID          string    `json:"Id,omitempty"`
	TxnDate     string    `json:"TxnDate"`
	TotalAmt    float64   `json:"TotalAmt"`
	Line        []Line    `json:"Line"`
	MetaData    Metadata  `json:"MetaData,omitempty"`
	SyncToken   string    `json:"SyncToken,omitempty"`
	CurrencyRef Reference `json:"CurrencyRef,omitempty"`
}
