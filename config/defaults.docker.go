//go:build docker

package config

import (
	"net/http"
	"time"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	errorMiddleware "github.com/wisdom-oss/microservice-middlewares/v5/error"
	securityMiddleware "github.com/wisdom-oss/microservice-middlewares/v5/security"

	"microservice/globals"
)

// This file contains default paths that are used inside the service to load
// resource files.
// The go toolchain automatically uses this file if no build tags have been
// set.
// This should be the case in local development and testing.
// To achieve the behavior in docker containers, supply the build tag "docker"
// and the Docker defaults are used

// Middlewares contains the middlewares used per default in this service.
// To disable single middlewares, please remove the line in which this array
// is used and add the middlewares that shall be used manually to the router
var Middlewares = []func(next http.Handler) http.Handler{
	httpLogger(),
	chiMiddleware.RequestID,
	chiMiddleware.RealIP,
	errorMiddleware.Handler,
	securityMiddleware.ValidateServiceJWT,
}

// EnvironmentFilePath contains the default file path under which the
// environment configuration file is stored
const EnvironmentFilePath = "./environment.json"

// QueryFilePath contains the default file path under which the
// sql queries are stored
const QueryFilePath = "./queries.sql"

// ListenAddress sets the host on which the microservice listens to incoming
// requests
const ListenAddress = "0.0.0.0"

func httpLogger() func(next http.Handler) http.Handler {
	l := httplog.NewLogger(globals.ServiceName)
	httplog.Configure(httplog.Options{
		JSON:            true,
		Concise:         true,
		TimeFieldFormat: time.RFC3339,
	})
	return httplog.Handler(l)

}
