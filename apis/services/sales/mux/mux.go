package mux

import (
	"net/http"

	"github.com/nico-phil/service/apis/services/sales/route/sys/checkapi"
)

func WebAPI() *http.ServeMux {
	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
