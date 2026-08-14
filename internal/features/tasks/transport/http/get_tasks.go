package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/zinovev-dm/golang-todoapp/internal/core/logger"
	core_http_request "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/response"
)

type GetTasksResponse []TasksDTOResponse

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userID, limit, offset, err := getUserIDLimitOffsetParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID/limit/offset query params")
		return
	}

	tasksDomains, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}

	response := GetTasksResponse(tasksDTOFromDomain(tasksDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDLimitOffsetParams(r *http.Request) (*int, *int, *int, error) {
	const (
		authorUserIDQueryParamKey = "user_id"
		limitQueryParamKey        = "limit"
		offsetQueryParamKey       = "offset"
	)
	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `%s` query param: %w", limitQueryParamKey, err)
	}
	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `%s` query param: %w", offsetQueryParamKey, err)
	}
	user_id, err := core_http_request.GetIntQueryParam(r, authorUserIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `%s` query param: %w", authorUserIDQueryParamKey, err)
	}

	return user_id, limit, offset, nil
}
