package quickbooks

// SetBaseURL allows to set the base URL (useful for switching between sandbox and production)
func (c *Client) SetBaseURL(baseURL string) {
	c.BaseURL = baseURL
}
