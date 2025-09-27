package mux

import (
	"os"

	"github.com/nico-phil/service/apis/services/sales/route/sys/checkapi"
	"github.com/nico-phil/service/app/api/mid"
	"github.com/nico-phil/service/fondation/logger"
	"github.com/nico-phil/service/fondation/web"
)

func WebAPI(log *logger.Logger, shutdonw chan os.Signal) *web.App {
	mux := web.New(shutdonw, mid.Logger(log))

	checkapi.Routes(mux)

	return mux
}
