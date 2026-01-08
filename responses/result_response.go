package responses

type ResultResponse[T any] struct {
	Success bool     `json:"success"`
	Status  int      `json:"status"`
	Data    *T       `json:"data,omitempty"`
	Error   *Error   `json:"error,omitempty"`
	Errors  *[]Error `json:"errors,omitempty"`
}

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func Success[T any](status int, data T) ResultResponse[T] {
	return ResultResponse[T]{
		Success: true,
		Status:  status,
		Data:    &data,
	}
}

func Failure[T any](status int, msg string) ResultResponse[T] {
	return ResultResponse[T]{
		Success: false,
		Status:  status,
		Error: &Error{
			Message: msg,
		},
	}
}

func Failures[T any](status int, errs map[string]string) ResultResponse[T] {
	errors := make([]Error, 0, len(errs))

	for field, msg := range errs {
		errors = append(errors, Error{
			Field:   field,
			Message: msg,
		})
	}
	return ResultResponse[T]{
		Success: false,
		Status:  status,
		Errors:  &errors,
	}
}
