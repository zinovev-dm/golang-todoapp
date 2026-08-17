package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/zinovev-dm/golang-todoapp/internal/core/logger"
	core_http_request "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/response"
)

type GetTaskResponse TasksDTOResponse

func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	id, err := core_http_request.GetIntQueryParam(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get query param id")
	}
	taskDomain, err := h.tasksService.GetTask(ctx, *id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}

	response := GetTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
