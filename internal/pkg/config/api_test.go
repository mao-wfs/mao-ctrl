package config

import (
	"testing"

	"github.com/mao-wfs/mao-ctrl/internal/pkg/testutil"
)

func TestLoadAPIConfig(t *testing.T) {
	pairs := map[string]string{
		"API_PORT": "3030",
	}

	reset := testutil.SetEnvs(t, pairs)
	defer reset()

	conf, err := LoadAPIConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.Port, int16(3030); got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestLoadAPIConfigPortDefault(t *testing.T) {
	conf, err := LoadAPIConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.Port, int16(3000); got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestAPIConfig_GetAddr(t *testing.T) {
	conf := &APIConfig{
		Port: 5000,
	}
	if got, want := conf.Addr(), ":5000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
