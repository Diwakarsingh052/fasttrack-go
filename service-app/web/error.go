package web

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Error is used to pass an error during the request through the
// application with web specific context.
// if someone is using this struct to construct an error than I can send the message saved in here to the end user otherwise I am going to send a generic message as InternalServerErro
type Error struct {
	Err    error
	Status int
}

func (err *Error) Error() string {
	return err.Err.Error()
}

func NewRequestError(err error, status int) error {
	return &Error{
		Err:    err,
		Status: status,
	}

}
