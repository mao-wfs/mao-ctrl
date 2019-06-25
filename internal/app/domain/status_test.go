package domain_test

import (
	"testing"

	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
)

const (
	waiting domain.Status = iota
	running
)

func TestNewStatus(t *testing.T) {
	s := domain.NewStatus()
	if got, want := *s, waiting; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestSetWaiting(t *testing.T) {
	s := domain.NewStatus()
	s.SetWaiting()
	if got, want := *s, waiting; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestSetRunning(t *testing.T) {
	s := domain.NewStatus()
	s.SetRunning()
	if got, want := *s, running; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestIsWaiting(t *testing.T) {
	s := domain.NewStatus()
	if got, want := s.IsWaiting(), true; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	s.SetRunning()
	if got, want := s.IsWaiting(), false; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestIsRunning(t *testing.T) {
	s := domain.NewStatus()
	if got, want := s.IsRunning(), false; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	s.SetRunning()
	if got, want := s.IsRunning(), true; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
