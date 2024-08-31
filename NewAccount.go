package quickbooks

// NewAccount creates a new instance of an Account struct with the provided details.
func NewAccount(name, accountType string) *Account {
	return &Account{
		Name:        name,
		AccountType: accountType,
		// You can add other fields if necessary.
	}
}
