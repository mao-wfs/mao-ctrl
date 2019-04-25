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
	Config config.Config
	*echo.Echo
}

func newRouter() (*Router, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	r := &Router{
		Config: conf,
		Echo:   echo.New(),
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
	apiConf := r.Config.GetAPIConfig()
	addr := apiConf.GetAddr()
	r.Logger.Fatal(r.Start(addr))
}

func (r *Router) initRouter() {
	r.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	conf, err := config.LoadConfig()
	if err != nil {
		r.Logger.Fatal(err)
	}

	devConf := conf.GetDeviceConfig()
	corrHan, err := correlator.NewHandler(devConf.GetCorrelatorConfig())
	if err != nil {
		r.Logger.Fatal(err)
	}
	fgHan, err := fg.NewHandler(devConf.GetFGConfig())
	if err != nil {
		r.Logger.Fatal(err)
	}
	wfsHan := device.NewWFSHandler(devConf, corrHan, fgHan)

	apiConf := conf.GetAPIConfig()
	api := r.Group(apiConf.GetRootEndPoint())
	{
		ctrl := controller.NewWFSController(conf, wfsHan)
		api.PUT("/start", handler.StartWFS(ctrl))
		api.PUT("/halt", handler.HaltWFS(ctrl))
	}
}
