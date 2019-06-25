package presenter

import (
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/output"
)

// WFSPresenter is the presenter of MAO-WFS services.
type WFSPresenter interface {
	ResponseError(code int, err error) output.Error
}

type wfsPresenter struct{}

// NewWFSPresenter returns the new presenter of MAO-WFS.
func NewWFSPresenter() WFSPresenter {
	return &wfsPresenter{}
}

func (p *wfsPresenter) ResponseError(code int, err error) output.Error {
	return NewError(code, err.Error())
}
