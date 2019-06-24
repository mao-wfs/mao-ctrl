package domain

const (
	waiting Status = iota
	running
)

// Status is the status of MAO-WFS.
type Status uint8

// NewStatus returns the new status.
// The initial value is 'waiting'.
func NewStatus() *Status {
	s := new(Status)
	*s = waiting
	return s
}

// SetWaiting sets the status 'waiting'.
func (s *Status) SetWaiting() { *s = waiting }

// SetRunning sets the status 'running'.
func (s *Status) SetRunning() { *s = running }

// IsWaiting checks whether the status is 'waiting'.
func (s *Status) IsWaiting() bool { return *s == waiting }

// IsRunning checks whether the status is 'running'.
func (s *Status) IsRunning() bool { return *s == running }
