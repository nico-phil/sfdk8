package checkapi

import (
	"github.com/nico-phil/service/fondation/web"
)

func Routes(app *web.App) {
	app.HandleFuncNoMidlleware("GET /liveness", liveness)
	app.HandleFuncNoMidlleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerr", testerr)
	app.HandleFunc("GET /testpanic", testpanic)
}
