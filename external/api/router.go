package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/mao-wfs/mao-ctrl/config"
	"github.com/mao-wfs/mao-ctrl/registry"
)

// Run runs the router of the API.
func Run() {
	r := newRouter()
	r.initialize()
	r.run()
}

// Router represents the router of MAO-WFS controller.
type Router struct {
	*echo.Echo
}

func newRouter() *Router {
	return &Router{
		Echo: echo.New(),
	}
}

func (r *Router) run() {
	conf, err := config.LoadAPIConfig()
	if err != nil {
		r.Logger.Fatal(err)
	}
	r.Logger.Fatal(r.Start(conf.GetAddr()))
}

func (r *Router) initialize() {
	r.Use(
		middleware.Logger(),
		middleware.Recover(),
	)
	ctn, err := registry.NewWFSContainer()
	if err != nil {
		r.Logger.Fatal(err)
	}

	ctrl := ctn.NewWFSController()
	r.GET("/start", func(c echo.Context) error {
		return ctrl.Start(c)
	})
	r.GET("/halt", func(c echo.Context) error {
		return ctrl.Halt(c)
	})

}
