package registry

import (
	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/controller"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device/correlator"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device/optswitch"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/presenter"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/input"
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/interactor"
)

type WFSContainer struct {
	correlator correlator.Handler
	optswitch  optswitch.Handler
}

func NewWFSContainer(corr correlator.Handler, sw optswitch.Handler) *WFSContainer {
	return &WFSContainer{
		correlator: corr,
		optswitch:  sw,
	}
}

func (c *WFSContainer) NewWFSController() controller.WFSController {
	return controller.NewWFSController(c.newWFSUsecase())
}

func (c *WFSContainer) newWFSUsecase() input.WFSInputPort {
	return interactor.NewWFSInteractor(domain.NewStatus(), c.newWFSHandler(), presenter.NewErrorPresenter())
}

func (c *WFSContainer) newWFSHandler() domain.Handler {
	return device.NewHandler(c.correlator, c.optswitch)
}
