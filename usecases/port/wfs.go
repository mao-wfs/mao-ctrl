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
	Halt(ctx context.Context, req *HaltWFSRequest) (*HaltWFSResponse, error)
}

// WFSOutputPort is the interface that describe the output port of MAO-WFS controller.
// TODO: Specify the presenter specifications.
type WFSOutputPort interface {
	// StartWFS starts MAO-WFS.
	Start(ctx context.Context) (*StartWFSResponse, error)
	// HaltWFS halts MAO-WFS.
	Halt(ctx context.Context) (*HaltWFSResponse, error)
}

// StartWFSRequest represents a request parameters to start MAO-WFS.
type StartWFSRequest struct {
	StartTime time.Time
	EndTime   time.Time
}

// GetStartTime returns the time to start MAO-WFS.
func (req *StartWFSRequest) GetStartTime() time.Time {
	return req.StartTime
}

// GetEndTime returns the time to halt MAO-WFS.
func (req *StartWFSRequest) GetEndTime() time.Time {
	return req.EndTime
}

// HaltWFSRequest represents a request parameters to halt MAO-WFS.
type HaltWFSRequest struct {
	HaltTime time.Time
}

// GetHaltTime returns the time to halt MAO-WFS.
func (req *HaltWFSRequest) GetHaltTime() time.Time {
	return req.HaltTime
}

// StartWFSResponse represents a response parameters of starting MAO-WFS.
// TODO: Specify the presenter specifications.
type StartWFSResponse struct{}

// HaltWFSResponse represents a response parameters of halting MAO-WFS.
// TODO: Specify the presenter specifications.
type HaltWFSResponse struct{}
