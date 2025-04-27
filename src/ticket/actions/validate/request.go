package validate

import "fmt"

type Request struct {
	Code string `json:"code"`
}

func (req *Request) Validate() error {
	if req.Code == "" {
		return fmt.Errorf("code is required")
	}
	return nil
}
