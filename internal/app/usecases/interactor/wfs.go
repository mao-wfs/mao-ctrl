package interactor

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/input"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/output"
)

const contextTimeout = 10 * time.Second

// WFSInteractor is the interactor of MAO-WFS services.
type WFSInteractor struct {
	status          *domain.Status
	handler         domain.Handler
	errorOutputPort output.ErrorOutputPort
}

// NewWFSInteractor returns a new interactor of MAO-WFS.
func NewWFSInteractor(s *domain.Status, h domain.Handler, errOpt output.ErrorOutputPort) input.WFSInputPort {
	return &WFSInteractor{
		status:          s,
		handler:         h,
		errorOutputPort: errOpt,
	}
}

// Start starts MAO-WFS.
func (i *WFSInteractor) Start(ctx context.Context) output.Error {
	subCtx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()

	if i.status.IsRunning() {
		err := xerrors.Errorf("MAO-WFS is already running")
		return i.errorOutputPort.BadRequest(subCtx, err)
	}
	if err := i.handler.Start(subCtx); err != nil {
		return i.errorOutputPort.InternalServerError(subCtx, err)
	}
	i.status.SetRunning()
	return nil
}

// Halt halts MAO-WFS.
func (i *WFSInteractor) Halt(ctx context.Context) output.Error {
	subCtx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()

	if i.status.IsWaiting() {
		err := xerrors.Errorf("MAO-WFS is not running")
		return i.errorOutputPort.BadRequest(subCtx, err)
	}
	if err := i.handler.Halt(subCtx); err != nil {
		return i.errorOutputPort.InternalServerError(subCtx, err)
	}
	i.status.SetWaiting()
	return nil
}
