package mid

import (
	"context"
	"net/http"

	"github.com/nico-phil/service/fondation/logger"
	"github.com/nico-phil/service/fondation/web"
)

func Logger(log *logger.Logger) web.MidHandler {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// logging start
			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			// logging end
			log.Info(ctx, "request complete", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
			return err
		}

		return h

	}

	return m

}
