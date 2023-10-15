package exception

type BadRequestError struct {
	Error error
}

func NewBadRequestError(error error) BadRequestError {
	return BadRequestError{
		Error: error,
	}
}
