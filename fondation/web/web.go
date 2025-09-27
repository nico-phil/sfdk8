package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type Handler func(context.Context, http.ResponseWriter, *http.Request) error

type App struct {
	*http.ServeMux
	shutdonw chan os.Signal
	mw       []MidHandler
}

func New(shutdonw chan os.Signal, mw ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdonw: shutdonw,
		mw:       mw,
	}
}

func (a *App) HandleFunc(pattern string, handler Handler, mw ...MidHandler) {

	handler = WrapMiddleware(mw, handler)
	handler = WrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		// any code here
		if err := handler(r.Context(), w, r); err != nil {
			fmt.Println(err)
		}

		// any code here
	}
	a.ServeMux.HandleFunc(pattern, h)

}
