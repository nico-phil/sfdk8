package web

import (
	"net/http"
	"os"
)

type App struct {
	*http.ServeMux
	shutdonw chan os.Signal
}

func New(shutdonw chan os.Signal) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdonw: shutdonw,
	}
}
