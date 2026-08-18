package core_http_middleware

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	core_logger "github.com/zinovev-dm/golang-todoapp/internal/core/logger"
	core_http_response "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/response"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-ID"

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}
			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)

		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			logger := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)
			ctx := core_logger.ToContext(r.Context(), logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(p, "during handle HTTP request got unexpected panic")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)
			startTime := time.Now()

			log.Debug(
				">>>> incoming HTTP request",
				zap.Time("time", startTime.UTC()),
				zap.String("HTTP Method", r.Method),
			)

			next.ServeHTTP(rw, r)
			log.Debug(
				"<<<< done HTTP request",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Since(startTime)),
			)
		})
	}
}

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := map[string]struct{}{
				"http://localhost:5050": {},
				"http://127.0.0.1:5050": {},
			}

			origin := r.Header.Get("Origin")

			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
