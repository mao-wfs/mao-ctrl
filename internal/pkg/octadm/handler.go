package octadm

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

// Handler is the handler of OCTAD-M (Elecs Industry Co., Ltd.)
type Handler interface {
	// Close closes the connection.
	Close() error

	// Start starts the correlation.
	Start(t time.Time, m CorrelationMode) error

	// Halt halts the correlation.
	Halt(t time.Time) error

	// SyncExternal synchronizes the internal sync signal with the external one.
	SyncExternal() error

	// Calibrate calibrates data transmission from ADC to FPGA.
	CalibrateDeMultiplexer(adcID int) error

	// SetDynamicRange sets the dynamic range of ADC.
	//
	// This sets the range and the offset voltage in mV.
	SetDynamicRange(adcID int, r, o float32) error

	// SetDelayOffset sets the delay offset of ADC.
	//
	// This sets it in sps (samples per second).
	SetDelayOffset(adcID int, offset uint) error

	// SetCorrelationScaling sets the scaling of X (Correlation) part.
	SetCorrelationScaling(threadID int, scale uint) error

	// SetRequantizationScaling sets the scaling of Y (Requantization) part.
	SetRequantizationScaling(threadID int, scale uint) error
	// SetIntegTime sets the integration time.
	//
	// Allowed values are 5 or 10 (ms).
	SetIntegTime(t IntegTime) error

	// SetMaskTime sets the time to mask the integration.
	//
	// This sets the time by FFT segment unit. (e.g. 2000 FFT segment is 2 ms).
	// Allowed values are 0 to 2000.
	SetMaskTime(t MaskTime) error

	// SetWindowFunc sets a window function.
	SetWindowFunc(w WindowFunc) error
}

// NewHandler returns a new handler of OCTAD-M (Elecs Industry Co., Ltd.).
func NewHandler(addr string, timeout time.Duration) (Handler, error) {
	c, err := newClient(addr, timeout)
	if err != nil {
		return nil, err
	}

	h := &handler{
		Client: c,
	}
	return h, nil
}

type handler struct {
	Client
}

func (h *handler) Start(t time.Time, m CorrelationMode) error {
	// TODO: Implement the validator of arguments
	cmd := fmt.Sprintf(
		"ctl_corstart=%04dy%03dd%02dh%02dm%02ds:0x%02x",
		t.Year(),
		t.YearDay(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		m,
	)
	return h.Exec(cmd)
}

func (h *handler) Halt(t time.Time) error {
	cmd := fmt.Sprintf(
		"ctl_corstop=%04dy%03dd%02dh%02dm%02ds",
		t.Year(),
		t.YearDay(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
	return h.Exec(cmd)
}

func (h *handler) SyncExternal() error {
	return h.Exec("ctl_sync")
}

func (h *handler) CalibrateDeMultiplexer(adcID int) error {
	// TODO: Implement the validator of arguments
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := make(chan error, 1)
	go func() {
		cmd := fmt.Sprintf("ctl_dmxcal%d", adcID)
		ch <- h.ExecContext(ctx, cmd)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		// BUG(scizorman): PingContext is not implemented yet.
		return xerrors.New("timeout")
	}
}

func (h *handler) SetDynamicRange(adcID int, volt, offset float32) error {
	// TODO: Implement the validator of arguments
	cmd := fmt.Sprintf("set_adc%d=%.1f:%.1f", adcID, volt, offset)
	return h.Exec(cmd)
}

func (h *handler) SetDelayOffset(adcID int, offset uint) error {
	// TODO: Implement the validator of arguments
	cmd := fmt.Sprintf("set_dlyoffset%01d=%d", adcID, offset)
	return h.Exec(cmd)
}

func (h *handler) SetCorrelationScaling(threadID int, scale uint) error {
	// TODO: Implement the validator of arguments
	cmd := fmt.Sprintf("set_scaling%d=%d", threadID, scale)
	return h.Exec(cmd)
}

func (h *handler) SetRequantizationScaling(threadID int, scale uint) error {
	// TODO: Implement the validator of arguments
	cmd := fmt.Sprintf("set_requantization%d=%d", threadID, scale)
	return h.Exec(cmd)
}

func (h *handler) SetIntegTime(t IntegTime) error {
	// TODO: Implement the validator of arguments
	cmd := fmt.Sprintf("set_iplen=%d", t)
	return h.Exec(cmd)
}

func (h *handler) SetMaskTime(t MaskTime) error {
	// BUG(scizorman): SetMaskTime is not implemented yet.
	return nil
}

func (h *handler) SetWindowFunc(w WindowFunc) error {
	// TODO: Implement the validator of arguments
	cmd := fmt.Sprintf("set_window=%s", w)
	return h.Exec(cmd)
}
