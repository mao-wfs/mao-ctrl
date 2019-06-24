package interactor

import (
	"context"

	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/input"
)

// WFSInteractor is the interactor of MAO-WFS services.
type WFSInteractor struct {
	status  *domain.Status
	handler domain.Handler
}

// NewWFSInteractor returns a new interactor of MAO-WFS.
func NewWFSInteractor(s *domain.Status, h domain.Handler) input.WFSInputPort {
	return &WFSInteractor{
		status:  s,
		handler: h,
	}
}

// Start starts MAO-WFS.
func (i *WFSInteractor) Start(ctx context.Context) error {
	if err := i.handler.Start(ctx); err != nil {
		return err
	}
	i.status.SetRunning()
	return nil
}

// Halt halts MAO-WFS.
func (i *WFSInteractor) Halt(ctx context.Context) error {
	if err := i.handler.Halt(ctx); err != nil {
		return err
	}
	i.status.SetWaiting()
	return nil
}

// IsRunning checks whether MAO-WFS is running.
func (i *WFSInteractor) IsRunning() bool {
	return i.status.IsRunning()
}

// IsWaiting checks whether MAO-WFS is waiting.
func (i *WFSInteractor) IsWaiting() bool {
	return i.status.IsWaiting()
}
