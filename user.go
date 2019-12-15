package detectlanguage

type UserStatusResponse struct {
	Date               string `json:"date,omitempty"`
	Requests           int    `json:"requests"`
	Bytes              int    `json:"bytes"`
	Plan               string `json:"plan"`
	PlanExpires        string `json:"plan_expires,omitempty"`
	DailyRequestsLimit int    `json:"daily_requests_limit"`
	DailyBytesLimit    int    `json:"daily_bytes_limit"`
	Status             string `json:"status"`
}

func (c *Client) UserStatus() (out *UserStatusResponse, err error) {
	err = c.get(nil, "user/status", &out)
	return
}
