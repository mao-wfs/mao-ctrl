package config

import (
	"os"
	"testing"
)

func TestLoadAPIConfig(t *testing.T) {
	pairs := map[string]string{
		"API_HOST": "127.0.0.1",
		"API_PORT": "3030",
	}

	reset := setEnvs(t, pairs)
	defer reset()

	conf, err := LoadAPIConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.Port, 3030; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestLoadAPIConfigPortDefault(t *testing.T) {
	pairs := map[string]string{
		"API_HOST": "127.0.0.1",
	}

	reset := setEnvs(t, pairs)
	defer reset()

	conf, err := LoadAPIConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.Port, 3000; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestAPIConfig_GetAddr(t *testing.T) {
	conf := &APIConfig{
		Host: "127.0.0.1",
		Port: 5000,
	}
	if got, want := conf.GetAddr(), "127.0.0.1:5000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func setEnvs(t *testing.T, envs map[string]string) func() {
	var resetFuncs []func()
	t.Helper()
	for k, v := range envs {
		r := setEnv(t, k, v)
		resetFuncs = append(resetFuncs, r)
	}
	return func() {
		for _, f := range resetFuncs {
			f()
		}
	}
}

func setEnv(t *testing.T, key, val string) func() {
	original := os.Getenv(key)
	if err := os.Setenv(key, val); err != nil {
		t.Fatal(err)
	}
	return func() {
		if original == "" {
			os.Unsetenv(key)
		} else {
			if err := os.Setenv(key, original); err != nil {
				t.Fatal(err)
			}
		}
	}
}
