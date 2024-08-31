package quickbooks

// Account represents a QuickBooks account.
type Account struct {
	AccountSubType                string    `json:"AccountSubType,omitempty"`
	AccountType                   string    `json:"AccountType"`
	Active                        bool      `json:"Active,omitempty"`
	Classification                string    `json:"Classification,omitempty"`
	CurrencyRef                   Reference `json:"CurrencyRef,omitempty"`
	CurrentBalance                float64   `json:"CurrentBalance,omitempty"`
	CurrentBalanceWithSubAccounts float64   `json:"CurrentBalanceWithSubAccounts,omitempty"`
	FullyQualifiedName            string    `json:"FullyQualifiedName,omitempty"`
	MetaData                      Metadata  `json:"Metadata"`
	ID                            string    `json:"Id,omitempty"`
	Name                          string    `json:"Name"`
	SubAccount                    bool      `json:"SubAccount,omitempty"`
	SyncToken                     int       `json:"SyncToken,omitempty"`
}
