package quickbooks

// Reference is used to refer to another entity such as Account, Customer, etc.
type Reference struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}
