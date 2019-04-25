package presenter

import (
	"context"

	"github.com/mao-wfs/mao-ctrl/config"
	"github.com/mao-wfs/mao-ctrl/usecases/port"
)

// WFSPresenter represents the presenter of MAO-WFS.
// TODO: Specify the presenter specifications.
type WFSPresenter struct {
	Config config.Config
}

// NewWFSPresenter returns a new presenter of MAO-WFS.
// TODO: Specify the presenter specifications.
func NewWFSPresenter(c config.Config) *WFSPresenter {
	return &WFSPresenter{
		Config: c,
	}
}

// Start starts MAO-WFS.
// TODO: Specify the presenter specifications.
func (p *WFSPresenter) Start(ctx context.Context) (*port.StartWFSResponse, error) {
	return nil, nil
}

// Halt halts MAO-WFS.
// TODO: Specify the presenter specifications.
func (p *WFSPresenter) Halt(ctx context.Context) (*port.HaltWFSResponse, error) {
	return nil, nil
}
