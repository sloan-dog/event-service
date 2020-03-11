package web

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	app := App{
		Router: gin.Default(),
	}
	return &app
}

func (a *App) Handle(verb, path string, handler gin.HandlerFunc) {
	// set up a route handler on gin router
	a.Router.Handle(verb, path, handler)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Router.ServeHTTP(w, r)
}

// ctxKey represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}
