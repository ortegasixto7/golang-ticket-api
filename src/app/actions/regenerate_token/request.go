package regenerate_token

import "fmt"

type Request struct {
	AppID string `json:"app_id" binding:"required"`
}

func (req *Request) Validate() error {
	if req.AppID == "" {
		return fmt.Errorf("app ID is required")
	}
	return nil
}
