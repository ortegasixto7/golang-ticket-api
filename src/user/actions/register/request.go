package register

type Request struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Request) Validate() error {
	if req.Name == "" {
		return ErrNameRequired
	}
	if req.Username == "" {
		return ErrUsernameRequired
	}
	if req.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

var (
	ErrNameRequired     = &RequestValidationError{Field: "name", Message: "name is required"}
	ErrUsernameRequired = &RequestValidationError{Field: "username", Message: "username is required"}
	ErrPasswordRequired = &RequestValidationError{Field: "password", Message: "password is required"}
)

type RequestValidationError struct {
	Field   string
	Message string
}

func (e *RequestValidationError) Error() string {
	return e.Message
}
