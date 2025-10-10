package handler

import "github.com/hokkung/release-management-service/internal/delivery/rest/model"

func handleAPIError(err error) *model.APIErrorResponse {
	return &model.APIErrorResponse{
		Error: model.APIError{
			Message: err.Error(),
		},
	}
}
