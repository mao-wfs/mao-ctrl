package interactor

import (
	"context"

	"github.com/mao-wfs/mao-ctrl/domain"
	"github.com/mao-wfs/mao-ctrl/usecases/port"
)

// WFSInteractor represents the interactor of MAO-WFS controller.
type WFSInteractor struct {
	Handler    domain.WFSHandler
	OutputPort port.WFSOutputPort
}

// NewWFSInteractor returns a new interactor of MAO-WFS controller.
func NewWFSInteractor(h domain.WFSHandler, op port.WFSOutputPort) *WFSInteractor {
	return &WFSInteractor{
		Handler:    h,
		OutputPort: op,
	}
}

// Start starts MAO-WFS.
func (i *WFSInteractor) Start(ctx context.Context, req *port.StartWFSRequest) (*port.StartWFSResponse, error) {
	st := req.GetStartTime()
	et := req.GetEndTime()
	if err := i.Handler.Start(ctx, st, et); err != nil {
		return nil, err
	}

	// TODO: Specify the presenter specifications.
	return i.OutputPort.Start(ctx)
}

// Halt halts MAO-WFS.
func (i *WFSInteractor) Halt(ctx context.Context, req *port.HaltWFSRequest) (*port.HaltWFSResponse, error) {
	ht := req.GetHaltTime()
	if err := i.Handler.Halt(ctx, ht); err != nil {
		return nil, err
	}

	// TODO: Specify the presenter specifications.
	return i.OutputPort.Halt(ctx)
}
