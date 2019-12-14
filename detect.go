package detectlanguage

type DetectRequest struct {
	Query string `json:"q"`
}

type DetectResponse struct {
	Data *DetectResponseData `json:"data"`
}

type DetectResponseData struct {
	Detections []*DetectionResult `json:"detections"`
}

type DetectionResult struct {
	Language   string  `json:"language"`
	Reliable   bool    `json:"isReliable"`
	Confidence float32 `json:"confidence"`
}

func (c *Client) Detect(in string) (out []*DetectionResult, err error) {
	var response DetectResponse
	err = c.post(nil, "detect", &DetectRequest{Query: in}, &response)

	if err != nil {
		return nil, err
	}

	return response.Data.Detections, err
}

func (c *Client) DetectCode(in string) (out string, err error) {
	detections, err := c.Detect(in)

	if err != nil {
		return "", err
	}

	if len(detections) == 0 {
		return "", &DetectionError{"Language not detected"}
	}

	return detections[0].Language, err
}
