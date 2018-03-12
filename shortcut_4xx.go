package response

import "net/http"

// BadRequest returns generic 4xx response.
func BadRequest() Response {
	return New().
		WithStatus(http.StatusBadRequest).
		WithError(
			NewError().
				WithTitle(http.StatusText(http.StatusBadRequest)).
				WithStatus(http.StatusBadRequest),
		)
}

// Unauthorized returns 401 Unauthorized response.
func Unauthorized() Response {
	return New().
		WithStatus(http.StatusUnauthorized).
		WithError(
			NewError().
				WithTitle(http.StatusText(http.StatusUnauthorized)).
				WithStatus(http.StatusUnauthorized),
		)
}

// Forbidden returns 403 Forbidden response.
func Forbidden() Response {
	return New().
		WithStatus(http.StatusForbidden).
		WithError(
			NewError().
				WithTitle(http.StatusText(http.StatusForbidden)).
				WithStatus(http.StatusForbidden),
		)
}

// NotFound returns 404 Not Found response.
func NotFound() Response {
	return New().
		WithStatus(http.StatusNotFound).
		WithError(
			NewError().
				WithTitle(http.StatusText(http.StatusNotFound)).
				WithStatus(http.StatusNotFound),
		)
}
