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
	sensT := req.GetSensingTime()
	stT, err := i.Handler.Start(ctx, sensT)
	if err != nil {
		return nil, err
	}
	return i.OutputPort.Start(ctx, stT, sensT)
}

// Halt halts MAO-WFS.
func (i *WFSInteractor) Halt(ctx context.Context) (*port.HaltWFSResponse, error) {
	hltT, err := i.Handler.Halt(ctx)
	if err != nil {
		return nil, err
	}
	return i.OutputPort.Halt(ctx, hltT)
}
