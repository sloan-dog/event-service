package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp(middleware ...gin.HandlerFunc) *App {
	app := App{
		Router: gin.New(),
	}
	app.Router.Use(middleware...) // set up root logger
	return &app
}

func (a *App) Handle(verb, path string, handlers ...gin.HandlerFunc) {

	a.Router.Handle(verb, path, handlers...)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Router.ServeHTTP(w, r)
}
