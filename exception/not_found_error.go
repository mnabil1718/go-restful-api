package exception

type NotFoundError struct {
	Message string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{error}
}
