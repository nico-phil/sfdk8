package all

import (
	"github.com/nico-phil/service/api/http/api/mux"
	"github.com/nico-phil/service/api/http/domain/checkapi"
	"github.com/nico-phil/service/fondation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
	})

	// testapi.Routes(app, testapi.Config{
	// 	Log:        cfg.Log,
	// 	AuthClient: cfg.AuthClient,
	// })

	// userapi.Routes(app, userapi.Config{
	// 	Log:        cfg.Log,
	// 	UserBus:    userBus,
	// 	AuthClient: cfg.AuthClient,
	// })

}
