package detectlanguage

// UserStatusResponse is the resource containing account status information
type AccountStatusResponse struct {
	Date               string `json:"date,omitempty"`
	Requests           int    `json:"requests"`
	Bytes              int    `json:"bytes"`
	Plan               string `json:"plan"`
	PlanExpires        string `json:"plan_expires,omitempty"`
	DailyRequestsLimit int    `json:"daily_requests_limit"`
	DailyBytesLimit    int    `json:"daily_bytes_limit"`
	Status             string `json:"status"`
}

// UserStatus fetches account status
func (c *Client) AccountStatus() (out *AccountStatusResponse, err error) {
	err = c.get(nil, "account/status", &out)
	return
}
