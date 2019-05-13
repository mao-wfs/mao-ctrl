package port

import (
	"context"
	"time"
)

// WFSInputPort is the interface that describe the input port of MAO-WFS controller.
type WFSInputPort interface {
	// Start starts MAO-WFS.
	Start(ctx context.Context, req *StartWFSRequest) (*StartWFSResponse, error)

	// HaltWFS halts MAO-WFS.
	Halt(ctx context.Context) (*HaltWFSResponse, error)
}

// WFSOutputPort is the interface that describe the output port of MAO-WFS controller.
// TODO: Specify the presenter specifications.
type WFSOutputPort interface {
	// StartWFS starts MAO-WFS.
	Start(ctx context.Context, stT time.Time, sensT time.Duration) (*StartWFSResponse, error)

	// HaltWFS halts MAO-WFS.
	Halt(ctx context.Context, hltT time.Time) (*HaltWFSResponse, error)
}

// StartWFSRequest represents a request parameters to start MAO-WFS.
type StartWFSRequest struct {
	SensingTime time.Duration
}

// GetSensingTime returns the time MAO-WFS senses.
func (req *StartWFSRequest) GetSensingTime() time.Duration {
	return req.SensingTime
}

// StartWFSResponse represents a response parameters after starting MAO-WFS.
type StartWFSResponse struct {
	StartTime time.Time
	EndTime   time.Time
}

// NewStartWFSResponse returns a response after starting MAO-WFS.
func NewStartWFSResponse(stT, edT time.Time) *StartWFSResponse {
	return &StartWFSResponse{
		StartTime: stT,
		EndTime:   edT,
	}
}

// HaltWFSResponse represents a response parameters after halting MAO-WFS.
type HaltWFSResponse struct {
	HaltTime time.Time
}

// NewHaltWFSResponse returns a response parameters after halting MAO-WFS.
func NewHaltWFSResponse(hltT time.Time) *HaltWFSResponse {
	return &HaltWFSResponse{
		HaltTime: hltT,
	}
}
