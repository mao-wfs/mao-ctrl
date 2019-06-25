package interactor

import (
	"context"
	"net/http"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/input"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/output"
)

var (
	errAlreadyRunning = xerrors.New("MAO-WFS is already running")
	errNotRunning     = xerrors.New("MAO-WFS is not running")
)

// WFSInteractor is the interactor of MAO-WFS services.
type WFSInteractor struct {
	status     *domain.Status
	handler    domain.Handler
	outputPort output.WFSOutputPort
}

// NewWFSInteractor returns a new interactor of MAO-WFS.
func NewWFSInteractor(s *domain.Status, h domain.Handler, opt output.WFSOutputPort) input.WFSInputPort {
	return &WFSInteractor{
		status:     s,
		handler:    h,
		outputPort: opt,
	}
}

// Start starts MAO-WFS.
func (i *WFSInteractor) Start(ctx context.Context) output.Error {
	if i.status.IsRunning() {
		return i.outputPort.ResponseError(http.StatusBadRequest, errAlreadyRunning)
	}
	if err := i.handler.Start(ctx); err != nil {
		return i.outputPort.ResponseError(http.StatusInternalServerError, err)
	}
	i.status.SetRunning()
	return nil
}

// Halt halts MAO-WFS.
func (i *WFSInteractor) Halt(ctx context.Context) output.Error {
	if i.status.IsWaiting() {
		return i.outputPort.ResponseError(http.StatusBadRequest, errNotRunning)
	}
	if err := i.handler.Halt(ctx); err != nil {
		return i.outputPort.ResponseError(http.StatusInternalServerError, err)
	}
	i.status.SetWaiting()
	return nil
}