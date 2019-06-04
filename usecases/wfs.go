package usecases

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/domain"
	"github.com/mao-wfs/mao-ctrl/usecases/input"
)

type wfsUsecase struct {
	handler domain.WFSHandler
}

func NewWFSUsecase(h domain.WFSHandler) input.WFSInputPort {
	return &wfsUsecase{
		handler: h,
	}
}

func (u *wfsUsecase) Start(ctx context.Context) error {
	if err := u.handler.Start(ctx); err != nil {
		return xerrors.Errorf("failed to start MAO-WFS: %w", err)
	}
	return nil
}

func (u *wfsUsecase) Halt(ctx context.Context) error {
	if err := u.handler.Halt(ctx); err != nil {
		return xerrors.Errorf("failed to halt MAO-WFS: %w", err)
	}
	return nil
}
