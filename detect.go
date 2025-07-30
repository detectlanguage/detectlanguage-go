package detectlanguage

import "context"

// DetectRequest contains language detection request params
type DetectRequest struct {
	Query string `json:"q"`
}

// DetectionResult is single language detection result
type DetectionResult struct {
	Language string  `json:"language"`
	Score    float64 `json:"score"`
}

// DetectBatchRequest contains batch language detection request params
type DetectBatchRequest struct {
	Query []string `json:"q"`
}

// Detect executes language detection for a single text
func (c *Client) Detect(ctx context.Context, in string) (out []*DetectionResult, err error) {
	var response []*DetectionResult
	err = c.post(ctx, "detect", &DetectRequest{Query: in}, &response)

	if err != nil {
		return nil, err
	}

	return response, err
}

// DetectCode executes language detection for a single text and returns detected language code
func (c *Client) DetectCode(ctx context.Context, in string) (out string, err error) {
	detections, err := c.Detect(ctx, in)

	if err != nil {
		return "", err
	}

	if len(detections) == 0 {
		return "", &DetectionError{"Language not detected"}
	}

	return detections[0].Language, err
}

// DetectBatch executes language detection with multiple texts.
// It is significantly faster than doing a separate request for each text indivdually.
func (c *Client) DetectBatch(ctx context.Context, in []string) (out [][]*DetectionResult, err error) {
	var response [][]*DetectionResult
	err = c.post(ctx, "detect-batch", &DetectBatchRequest{Query: in}, &response)

	if err != nil {
		return nil, err
	}

	return response, err
}
