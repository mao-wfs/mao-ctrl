package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/mao-wfs/mao-ctrl/adapters/controller"
	"github.com/mao-wfs/mao-ctrl/adapters/gateway/device"
	"github.com/mao-wfs/mao-ctrl/config"
	"github.com/mao-wfs/mao-ctrl/external/device/correlator"
	"github.com/mao-wfs/mao-ctrl/external/device/fg"
	"github.com/mao-wfs/mao-ctrl/external/server/handler"
)

// Router is the router of MAO-WFS controller.
type Router struct {
	*echo.Echo
}

func newRouter() (*Router, error) {
	r := &Router{
		Echo: echo.New(),
	}
	return r, nil
}

// Run run the server to control MAO-WFS.
func Run() {
	r, err := newRouter()
	if err != nil {
		panic(err)
	}
	r.initRouter()
	r.run()
}

func (r *Router) run() {
	conf, err := config.LoadAPIConfig()
	if err != nil {
		r.Logger.Fatal(err)
	}
	r.Logger.Fatal(r.Start(conf.GetAddr()))
}

func (r *Router) initRouter() {
	r.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	corrHan, err := correlator.NewHandler()
	if err != nil {
		r.Logger.Fatal(err)
	}
	fgHan, err := fg.NewHandler()
	if err != nil {
		r.Logger.Fatal(err)
	}
	wfsHan := device.NewWFSHandler(corrHan, fgHan)

	api := r.Group("api")
	{
		ctrl := controller.NewWFSController(wfsHan)
		api.PUT("/start", handler.StartWFS(ctrl))
		api.PUT("/halt", handler.HaltWFS(ctrl))
	}
}
