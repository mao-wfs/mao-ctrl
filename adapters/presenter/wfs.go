package presenter

import (
	"context"
	"time"

	"github.com/mao-wfs/mao-ctrl/usecases/port"
)

// WFSPresenter represents the presenter of MAO-WFS.
type WFSPresenter struct {
	port.WFSOutputPort
}

// NewWFSPresenter returns a new presenter of MAO-WFS.
// TODO: Specify the presenter specifications.
func NewWFSPresenter() *WFSPresenter {
	return &WFSPresenter{}
}

// Start starts MAO-WFS.
func (p *WFSPresenter) Start(ctx context.Context, stT time.Time, sensT time.Duration) (*port.StartWFSResponse, error) {
	edT := stT.Add(sensT)
	return port.NewStartWFSResponse(stT, edT), nil
}

// Halt halts MAO-WFS.
func (p *WFSPresenter) Halt(ctx context.Context, hltT time.Time) (*port.HaltWFSResponse, error) {
	return port.NewHaltWFSResponse(hltT), nil
}
