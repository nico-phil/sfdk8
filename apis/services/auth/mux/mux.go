package mux

import (
	"os"

	"github.com/nico-phil/service/apis/services/api/mid"
	"github.com/nico-phil/service/apis/services/auth/route/checkapi"
	"github.com/nico-phil/service/business/api/auth"
	"github.com/nico-phil/service/fondation/logger"
	"github.com/nico-phil/service/fondation/web"
)

func WebAPI(log *logger.Logger, auth auth.Auth, shutdonw chan os.Signal) *web.App {
	mux := web.New(shutdonw, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(mux)

	return mux
}
