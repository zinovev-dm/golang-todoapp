package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/zinovev-dm/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/zinovev-dm/golang-todoapp/internal/core/transport/http/server"
	users_transport_http "github.com/zinovev-dm/golang-todoapp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	fmt.Println("Hello, Todo app")
	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("starting todo application")

	userTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)

	usersRouters := userTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersionV1)
	apiVersionRouter.RegisterRoutes(usersRouters...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		*logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterAPIRoutes(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP Server run error", zap.Error(err))
	}
}
