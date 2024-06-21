package rest

import "net/http"

type ErrorPayload struct {
	Type  string `json:"type" example:"srn:error:some_error"`
	Title string `json:"title,omitempty" example:"Message for some error"`
}

var (
	errInternalServerError = ErrorPayload{Type: "error:internal_server_error", Title: "Internal Server Error"}
	errBadRequest          = ErrorPayload{Type: "error:bad_request", Title: "Bad Request"}
)

type Response struct {
	Body       any
	Error      error
	header     http.Header
	StatusCode int
}

func Created(body any) Response {
	return Response{
		StatusCode: http.StatusCreated,
		Body:       body,
	}
}

func OK(body any) Response {
	return Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}
}

func BadRequest(body any, err error) Response {
	return Response{
		StatusCode: http.StatusBadRequest,
		Body:       errBadRequest,
		Error:      err,
	}
}

func InternalServerError(body any, err error) Response {
	return Response{
		StatusCode: http.StatusInternalServerError,
		Body:       errInternalServerError,
		Error:      err,
	}
}
