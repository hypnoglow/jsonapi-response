// Package response provides Response type for use with APIs that conform
// http://jsonapi.org specification.
package response

import (
	"net/http"

	// Upstream repo https://github.com/google/jsonapi is out of date.
	// At least this issue https://github.com/google/jsonapi/issues/123 has to be
	// fixed before we can revert to upstream.
	"github.com/hypnoglow/jsonapi"
)

// Response is JSON:API response.
type Response struct {
	// statusCode is the response HTTP status code.
	statusCode int

	// data is the document's “primary data”
	data       interface{}

	// errors is an array of error objects
	errors []*jsonapi.ErrorObject
}

// New returns a new Response, which status code is 200 by default.
func New() Response {
	return Response{statusCode: http.StatusOK}
}

// WithStatus returns response with the specific status code.
func (r Response) WithStatus(code int) Response {
	r.statusCode = code
	return r
}

// WithData returns response with the specific data.
func (r Response) WithData(data interface{}) Response {
	r.data = data
	return r
}

// WithError returns response with the specific error.
func (r Response) WithError(err Error) Response {
	r.errors = append(r.errors, err.origin)
	return r
}

// Write writes the response to w.
func (r Response) Write(w http.ResponseWriter) error {
	if r.statusCode == http.StatusNoContent {
		w.WriteHeader(r.statusCode)
		return nil
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)

	w.WriteHeader(r.statusCode)

	var err error
	switch {
	case r.statusCode >= 200 && r.statusCode <= 299:
		err = jsonapi.MarshalPayload(w, r.data)
	default:
		err = jsonapi.MarshalErrors(w, r.errors)
	}

	return err
}

// MustWrite writes the response to w or panics if any error occurs.
func (r Response) MustWrite(w http.ResponseWriter) {
	if err := r.Write(w); err != nil {
		panic(err)
	}
}
