package controller

import (
	"context"
	"net/http"

	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/input"
)

type successResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newSuccessResponse(msg string) *successResponse {
	return &successResponse{
		Code:    http.StatusOK,
		Message: msg,
	}
}

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

	if err := ctrl.inputPort.Start(ctx); err != nil {
		c.JSON(err.StatusCode(), err)
		return err
	}
	return c.JSON(http.StatusOK, newSuccessResponse("MAO-WFS starts!"))
}

func (ctrl *wfsController) Halt(c Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := ctrl.inputPort.Halt(ctx); err != nil {
		c.JSON(err.StatusCode(), err)
		return err
	}
	return c.JSON(http.StatusOK, newSuccessResponse("MAO-WFS stopped"))
}
