package exception

type LoginError struct {
	Error error
}

func NewLoginError(error error) LoginError {
	return LoginError{
		Error: error,
	}
}
