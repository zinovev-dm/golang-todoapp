package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/zinovev-dm/golang-todoapp/internal/core/domain"
	core_logger "github.com/zinovev-dm/golang-todoapp/internal/core/logger"
	core_http_request "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/response"
	core_http_types "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/types"
)

type PathTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}
type PathTaskResponse TasksDTOResponse

func (r *PathTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("Title can't be NULL")
		}

		titleLen := len([]rune(*r.Title.Value))

		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("Title lenght must been between 1 and 100")
		}
	}
	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLen := len([]rune(*r.Description.Value))

			if descriptionLen < 1 || descriptionLen > 1000 {
				return fmt.Errorf("Description lenght must been between 1 and 1000")
			}
		}
	}
	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("Completed can't be NULL")
		}
	}

	return nil
}

func (h *TasksHTTPHandler) PathTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	taskID, err := core_http_request.GetIntQueryParam(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get query param id")
	}

	var request PathTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}
	taskPath := taskPathFromRequest(request)
	taskDomain, err := h.tasksService.PathTask(ctx, *taskID, taskPath)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to path task")

		return
	}
	respone := PathTaskResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JSONResponse(respone, http.StatusOK)
}

func taskPathFromRequest(request PathTaskRequest) domain.TaskPath {
	return domain.NewTaskPath(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}
