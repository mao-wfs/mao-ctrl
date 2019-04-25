package handler

import (
	"github.com/labstack/echo"

	"github.com/mao-wfs/mao-ctrl/adapters/controller"
)

func StartWFS(ctrl *controller.WFSController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctrl.Start(c)
	}
}

func HaltWFS(ctrl *controller.WFSController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctrl.Halt(c)
	}
}
