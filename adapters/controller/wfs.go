package controller

import (
	"context"
	"net/http"

	"github.com/mao-wfs/mao-ctrl/adapters/gateway/device"
	"github.com/mao-wfs/mao-ctrl/adapters/presenter"
	"github.com/mao-wfs/mao-ctrl/usecases/interactor"
	"github.com/mao-wfs/mao-ctrl/usecases/port"
)

// WFSController represents the controller of MAO-WFS.
type WFSController struct {
	InputPort port.WFSInputPort
}

// NewWFSController returns a new controller of MAO-WFS.
func NewWFSController(h *device.WFSHandler) *WFSController {
	p := presenter.NewWFSPresenter()
	return &WFSController{
		InputPort: interactor.NewWFSInteractor(h, p),
	}
}

// Start starts MAO-WFS.
func (ctrl *WFSController) Start(c Context) error {
	req := new(port.StartWFSRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := ctrl.InputPort.Start(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// Halt halts MAO-WFS.
func (ctrl *WFSController) Halt(c Context) error {
	req := new(port.HaltWFSRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := ctrl.InputPort.Halt(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}
