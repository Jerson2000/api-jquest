package responses

import (
	"math"
)

type ResultResponse[T any] struct {
	Success bool     `json:"success"`
	Status  int      `json:"status"`
	Data    *T       `json:"data,omitempty"`
	Error   *Error   `json:"error,omitempty"`
	Errors  *[]Error `json:"errors,omitempty"`
}

type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	TotalCount int64 `json:"totalCount"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"totalPages"`
	HasNext    bool  `json:"hasNext"`
	HasPrev    bool  `json:"hasPrev"`
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

func PaginatedSuccess[T any](status int, data []T, count int64, page, limit int) ResultResponse[PaginatedResponse[T]] {
	totalPages := int(math.Ceil(float64(count) / float64(limit)))
	return ResultResponse[PaginatedResponse[T]]{
		Success: true,
		Status:  status,
		Data: &PaginatedResponse[T]{
			Data:       data,
			TotalCount: count,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
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
