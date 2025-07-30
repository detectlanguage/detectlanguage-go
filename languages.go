package detectlanguage

import "context"

// Language is the resource representing language
type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Languages retrieves the list of supported languages
func (c *Client) Languages(ctx context.Context) (out []*Language, err error) {
	err = c.get(ctx, "languages", &out)
	return
}
