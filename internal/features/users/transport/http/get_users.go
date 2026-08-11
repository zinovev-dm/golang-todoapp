package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/zinovev-dm/golang-todoapp/internal/core/logger"
	core_http_request "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/response"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset param")
		return
	}

	users, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := GetUsersResponse(usersDTOFromDomains(users))

	responseHandler.JSONResponse(response, http.StatusOK)

}

func getLimitOffsetParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get `%s` query param: %w", limitQueryParamKey, err)
	}
	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get `%s` query param: %w", offsetQueryParamKey, err)
	}

	return limit, offset, nil
}
