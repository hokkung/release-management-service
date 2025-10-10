package model

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type APIDataResponse[T any] struct {
	Data T `json:"data"`
}

type APIResponse struct {
}

type APIErrorResponse struct {
	Error APIError `json:"error"`
}

type APIError struct {
	Message string `json:"message"`
}
