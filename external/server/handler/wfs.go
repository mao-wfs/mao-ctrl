package handler

import (
	"github.com/labstack/echo"

	"github.com/mao-wfs/mao-ctrl/adapters/controller"
)

// StartWFS starts MAO-WFS.
// It is the handler function of "Echo"
func StartWFS(ctrl *controller.WFSController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctrl.Start(c)
	}
}

// HaltWFS halts MAO-WFS.
// It is the handler function of "Echo"
func HaltWFS(ctrl *controller.WFSController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctrl.Halt(c)
	}
}
