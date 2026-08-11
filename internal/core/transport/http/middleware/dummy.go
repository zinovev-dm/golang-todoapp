package core_http_middleware

import (
	"fmt"
	"net/http"

	core_logger "github.com/zinovev-dm/golang-todoapp/internal/core/logger"
)

func Dummy(s string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)

			logger.Debug(fmt.Sprintf("-> before: %s", s))

			next.ServeHTTP(w, r)

			logger.Debug(fmt.Sprintf("<- after: %s", s))
		})
	}
}
