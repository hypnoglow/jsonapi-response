package response

import (
	"strconv"

	"github.com/hypnoglow/jsonapi" // See response.go for details.
)

// Error is JSON:API response error.
type Error struct {
	origin *jsonapi.ErrorObject
}

// NewError returns a new Error.
func NewError() Error {
	return Error{origin: &jsonapi.ErrorObject{}}
}

// WithTitle returns error with a short, human-readable summary of the problem.
func (e Error) WithTitle(title string) Error {
	e.origin.Title = title
	return e
}

// WithStatus returns error with the HTTP status code applicable to the problem.
func (e Error) WithStatus(status int) Error {
	e.origin.Status = strconv.Itoa(status)
	return e
}
