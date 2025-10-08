package web

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Handler func(context.Context, http.ResponseWriter, *http.Request) error

type Logger func(ctx context.Context, msg string, v ...any)

type App struct {
	*http.ServeMux
	log Logger
	mw  []MidHandler
}

func NewApp(log Logger, mw ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		log:      log,
		mw:       mw,
	}

}

func (a *App) HandleFunc(pattern string, handler Handler, mid ...MidHandler) {

	// handler = WrapMiddleware(mw, handler)
	// handler = WrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		v := Values{
			TraceID: uuid.NewString(),
			Now:     time.Now().UTC(),
		}
		ctx := setValues(r.Context(), &v)

		if err := handler(ctx, w, r); err != nil {
			a.log(ctx, "web", "ERROR", err)
			return
		}

	}
	a.ServeMux.HandleFunc(pattern, h)

}

// func (a *App) HandleFuncNoMidlleware(pattern string, handler Handler, mw ...MidHandler) {

// 	h := func(w http.ResponseWriter, r *http.Request) {
// 		v := Values{
// 			TraceID: uuid.NewString(),
// 			Now:     time.Now().UTC(),
// 		}
// 		ctx := setValues(r.Context(), &v)

// 		if err := handler(ctx, w, r); err != nil {
// 			fmt.Println(err)
// 		}

// 	}
// 	a.ServeMux.HandleFunc(pattern, h)

// }
