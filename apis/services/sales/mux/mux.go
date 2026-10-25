// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/rif-7/service/apis/services/api/mid"
	"github.com/rif-7/service/apis/services/sales/route/sys/checkapi"
	"github.com/rif-7/service/foundation/logger"
	"github.com/rif-7/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))

	checkapi.Routes(mux)

	return mux
}
