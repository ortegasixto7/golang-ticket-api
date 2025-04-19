package signup

import "github.com/ortegasixto7/golang-ticket/src/integration"

func Execute(req *Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	integration := integration.Integration{
		Name:        req.Name,
		Description: req.Description,
	}
	integration.GenerateToken()

	response := &Response{
		ID:          integration.ID,
		Name:        integration.Name,
		Description: integration.Description,
		AppToken:    integration.AppToken,
	}

	return response, nil
}
