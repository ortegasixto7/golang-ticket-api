package signup

import "fmt"

type Request struct {
	Name        string
	Description string
}

func (req *Request) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if req.Description == "" {
		return fmt.Errorf("description is required")
	}
	return nil
}
