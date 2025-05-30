package login

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Request) Validate() error {
	if req.Username == "" {
		return ErrUsernameRequired
	}
	if req.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

var (
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
