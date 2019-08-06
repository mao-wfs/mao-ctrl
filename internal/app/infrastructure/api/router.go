package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/mao-wfs/mao-ctrl/internal/app/configs"
	"github.com/mao-wfs/mao-ctrl/internal/app/infrastructure/device/correlator"
	"github.com/mao-wfs/mao-ctrl/internal/app/infrastructure/device/optswitch"
	"github.com/mao-wfs/mao-ctrl/internal/app/registry"
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
	conf, err := configs.LoadAPIConfig()
	if err != nil {
		r.Logger.Fatal(err)
	}
	r.Logger.Fatal(r.Start(conf.Addr()))
}

func (r *Router) initialize() {
	r.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	corr, err := correlator.NewHandler()
	if err != nil {
		r.Logger.Fatal(err)
	}

	sw, err := optswitch.NewHandler()
	if err != nil {
		r.Logger.Fatal(err)
	}

	ctn := registry.NewWFSContainer(corr, sw)

	ctrl := ctn.NewWFSController()
	r.GET("/start", func(c echo.Context) error {
		return ctrl.Start(c)
	})
	r.GET("/halt", func(c echo.Context) error {
		return ctrl.Halt(c)
	})

}
