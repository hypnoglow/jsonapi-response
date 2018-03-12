package response

import "net/http"

// OK returns 200 OK response.
func OK(data interface{}) Response {
	return New().
		WithStatus(http.StatusOK).
		WithData(data)
}

// Created returns 201 Created response..
func Created(data interface{}) Response {
	return New().
		WithStatus(http.StatusCreated).
		WithData(data)
}
