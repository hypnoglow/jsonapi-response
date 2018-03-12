package response

import "net/http"

// InternalServerError returns 500 Internal Server Error response.
func InternalServerError() Response {
	return New().
		WithStatus(http.StatusInternalServerError).
		WithError(
			NewError().
				WithTitle(http.StatusText(http.StatusInternalServerError)).
				WithStatus(http.StatusInternalServerError),
		)
}
