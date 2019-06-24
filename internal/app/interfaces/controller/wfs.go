package controller

import (
	"context"
	"net/http"

	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/input"
)

// WFSController is the interface that describe the controller of MAO-WFS.
type WFSController interface {
	// Start starts MAO-WFS.
	Start(c Context) error

	// Halt halts MAO-WFS.
	Halt(c Context) error
}

// WFSController represents the controller of MAO-WFS.
type wfsController struct {
	inputPort input.WFSInputPort
}

// NewWFSController returns a new controller of MAO-WFS.
func NewWFSController(ipt input.WFSInputPort) WFSController {
	return &wfsController{
		inputPort: ipt,
	}
}

func (ctrl *wfsController) Start(c Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if ctrl.inputPort.IsRunning() {
		apiErr := NewError(http.StatusBadRequest, "already running")
		return c.JSON(apiErr.Code(), apiErr)
	}

	if err := ctrl.inputPort.Start(ctx); err != nil {
		apiErr := NewError(http.StatusInternalServerError, err.Error())
		return c.JSON(apiErr.Code(), apiErr)
	}
	return c.JSON(http.StatusOK, "MAO-WFS started!")
}

func (ctrl *wfsController) Halt(c Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if ctrl.inputPort.IsWaiting() {
		apiErr := NewError(http.StatusBadRequest, "not running")
		return c.JSON(apiErr.Code(), apiErr)
	}

	if err := ctrl.inputPort.Halt(ctx); err != nil {
		apiErr := NewError(http.StatusInternalServerError, err.Error())
		return c.JSON(apiErr.Code(), apiErr)
	}
	return c.JSON(http.StatusOK, "MAO-WFS stoped!")
}
