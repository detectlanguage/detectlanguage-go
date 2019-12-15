package detectlanguage

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (c *Client) Languages() (out []*Language, err error) {
	err = c.get(nil, "languages", &out)
	return
}
