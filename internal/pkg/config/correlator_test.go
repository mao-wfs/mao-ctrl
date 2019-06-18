package config

import (
	"testing"

	"github.com/mao-wfs/mao-ctrl/internal/pkg/testutil"
)

func TestLoadCorrelatorConfig(t *testing.T) {
	pairs := map[string]string{
		"CORRELATOR_HOST": "127.0.0.1",
		"CORRELATOR_PORT": "5000",
	}

	reset := testutil.SetEnvs(t, pairs)
	defer reset()

	conf, err := LoadCorrelatorConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.Port, int16(5000); got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestCorrelatorConfig_GetAddr(t *testing.T) {
	conf := &CorrelatorConfig{
		Host: "127.0.0.1",
		Port: 5000,
	}
	if got, want := conf.Addr(), "127.0.0.1:5000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
