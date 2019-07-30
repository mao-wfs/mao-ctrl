package optswitch

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/scizorman/go-scpi"

	"github.com/mao-wfs/mao-ctrl/internal/app/configs"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device/optswitch"
)

const dialTimeout = 5 * time.Second

const initTimeout = 3 * time.Second

const (
	numRepeatPattern = 100
	integrationTime  = 0.005
)

type handler struct {
	config *configs.OptSwitchConfig
	pg     *scpi.Handler
	fg     *scpi.Handler
}

// NewHandler returns a new handler of the optical switch.
func NewHandler() (optswitch.Handler, error) {
	conf, err := configs.LoadOptSwitchConfig()
	if err != nil {
		return nil, err
	}
	pgClt, err := scpi.NewClient("tcp", conf.PGAddr(), dialTimeout)
	if err != nil {
		return nil, err
	}
	fgClt, err := scpi.NewClient("tcp", conf.FGAddr(), dialTimeout)
	if err != nil {
		return nil, err
	}

	h := &handler{
		config: conf,
		pg:     scpi.NewHandler(pgClt),
		fg:     scpi.NewHandler(fgClt),
	}
	return h, nil
}

// Start implements the Handler Start method.
func (h *handler) Start(ctx context.Context) error {
	if err := h.fg.Exec("INIT1:CONT ON"); err != nil {
		return err
	}
	if err := h.pg.Exec("OUTP ON"); err != nil {
		return err
	}
	time.Sleep(950 * time.Millisecond)

	if err := h.fg.BulkExec("INIT1:CONT OFF", "OUTP1 OFF"); err != nil {
		return err
	}
	return nil
}

// Halt implements the Handler Halt method.
func (h *handler) Halt(ctx context.Context) error {
	return h.pg.Exec("OUTP OFF")
}

// Initialize implements the Handler Initialize method.
func (h *handler) Initialize(ctx context.Context) error {
	if err := h.initPG(initTimeout); err != nil {
		return err
	}
	return h.initFG(initTimeout)
}

func (h *handler) initPG(d time.Duration) error {
	freq := numRepeatPattern / integrationTime
	ptn := h.tokenizeOrder()
	cmds := []string{
		"DIG:PATT ON",
		"DIG:PATT:CLOC POS",
		"DIG:PATT:REP ON",
		"DIG:PATT:TRIG:SLOP POS",
		"DIG:PATT:TRIG:SOUR EXT",

		"FUNC:PATT VOLATILE",
		fmt.Sprintf("DATA:PATTERN VOLATILE, %s", ptn),
		fmt.Sprintf("DIG:PATT:FREQ %f", freq),
		"DIG:PATT:FREQ 1000",
	}
	if err := h.pg.BulkExec(cmds...); err != nil {
		return err
	}

	return h.pg.WaitForComplete(d)
}

func (h *handler) tokenizeOrder() string {
	ptns := h.config.Order()
	arr := make([]string, numRepeatPattern*len(ptns))
	for i, p := range ptns {
		strP := strconv.Itoa(int(p))
		for j := 0; j < 100; j++ {
			arr[i*numRepeatPattern+j] = strP
		}
	}
	return strings.Join(arr, ", ")
}

func (h *handler) initFG(d time.Duration) error {
	cmds := []string{
		"SOUR1:FUNC:PULS",
		"SOUR1:FUNC:PULS:DCYC 5.0",
		"SOUR1:FUNC:PULS:HOLD DCYC",
		"SOUR1:FUNC:PULS:PER 1.",
		"SOUR1:FUNC:PULS:TRAN:BOTH 1.0E-8",

		"SOUR1:VOLT:HIGH 3.3",
		"SOUR1:VOLT:LOW 0.0",
		"SOUR1:VOLT:LIM:HIGH 3.4",
		"SOUR1:VOLT:LIM:LOW -1.0",
		"SOUR1:VOLT:LIM:STAT ON",

		"TRIG1:DEL 2.0E-2",
		"TRIG1:SLOP POS",
		"TRIG1:SOUR EXT",

		"SOUR1:BURS:MODE TRIG",
		"SOUR1:BURS:NCYC 1",
		"SOUR1:BURS:STAT ON",

		"INIT1:CONT OFF",
		"OUTP1 ON",
	}
	if err := h.fg.BulkExec(cmds...); err != nil {
		return err
	}

	return h.fg.WaitForComplete(d)
}
