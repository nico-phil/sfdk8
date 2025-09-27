package mux

import (
	"os"

	"github.com/nico-phil/service/apis/services/sales/route/sys/checkapi"
	"github.com/nico-phil/service/fondation/web"
)

func WebAPI(shutdonw chan os.Signal) *web.App {
	mux := web.New(shutdonw)

	checkapi.Routes(mux)

	return mux
}
