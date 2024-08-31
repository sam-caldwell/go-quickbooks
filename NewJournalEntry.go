package quickbooks

// NewJournalEntry creates a new instance of a JournalEntry with the given transaction date.
func NewJournalEntry(txnDate string) *Journal {
	return &Journal{
		TxnDate: txnDate,
		Line:    []Line{},
	}
}

// NewLine adds a new line item to the JournalEntry and returns a pointer to the updated JournalEntry.
func (je *Journal) NewLine(amount float64, description, postingType, accountRefValue, accountRefName string) *Journal {
	line := Line{
		Amount:      amount,
		Description: description,
		JournalLineDetail: LineDetail{
			PostingType: postingType,
			AccountRef: Reference{
				Value: accountRefValue,
				Name:  accountRefName,
			},
		},
	}
	je.Line = append(je.Line, line)
	je.TotalAmt += amount // Adjust the total amount of the Journal Entry.
	return je
}
