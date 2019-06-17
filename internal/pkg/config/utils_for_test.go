package config

import (
	"os"
	"testing"
)

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
