package registry

import (
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/controller"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device"
	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
	"github.com/mao-wfs/mao-ctrl/internal/app/infrastructure/device/correlator"
	"github.com/mao-wfs/mao-ctrl/internal/app/infrastructure/device/fg"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/input"
)

type WFSContainer struct {
	correlator device.CorrelatorHandler
	fg         device.FGHandler
}

func NewWFSContainer() (*WFSContainer, error) {
	corr, err := correlator.NewHandler()
	if err != nil {
		return nil, err
	}
	fg, err := fg.NewHandler()
	if err != nil {
		return nil, err
	}

	c := &WFSContainer{
		correlator: corr,
		fg:         fg,
	}
	return c, nil
}

func (c *WFSContainer) NewWFSController() controller.WFSController {
	return controller.NewWFSController(c.newWFSUsecase())
}

func (c *WFSContainer) newWFSUsecase() input.WFSInputPort {
	return usecases.NewWFSUsecase(c.newWFSHandler())
}

func (c *WFSContainer) newWFSHandler() domain.WFSHandler {
	return device.NewWFSHandler(c.correlator, c.fg)
}
