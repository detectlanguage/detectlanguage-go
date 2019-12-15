package detectlanguage

// Language is the resource representing language
type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Languages retrieves the list of supported languages
func (c *Client) Languages() (out []*Language, err error) {
	err = c.get(nil, "languages", &out)
	return
}
