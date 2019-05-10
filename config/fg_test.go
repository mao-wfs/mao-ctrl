package config

import (
	"reflect"
	"testing"
)

func TestLoadFGConfig(t *testing.T) {
	pairs := map[string]string{
		"FG_HOST":  "127.0.0.1",
		"FG_PORT":  "5000",
		"FG_ORDER": "10,9,13,8,0",
	}

	reset := setEnvs(t, pairs)
	defer reset()

	conf, err := LoadFGConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.Port, 5000; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if got, want := conf.Order, []int{10, 9, 13, 8, 0}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestLoadFGConfigOrderDefault(t *testing.T) {
	pairs := map[string]string{
		"FG_HOST": "127.0.0.1",
		"FG_PORT": "5000",
	}

	reset := setEnvs(t, pairs)
	defer reset()

	conf, err := LoadFGConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.Port, 5000; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if got, want := conf.Order, []int{10, 9, 13, 8, 0, 80, 16, 32}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestFGConfig_GetAddr(t *testing.T) {
	conf := &FGConfig{
		Host: "127.0.0.1",
		Port: 5000,
	}
	if got, want := conf.GetAddr(), "127.0.0.1:5000"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestFGConfig_GetOrder(t *testing.T) {
	conf := &FGConfig{
		Order: []int{10, 9, 13, 8, 0, 80, 16, 32},
	}
	if got, want := conf.GetOrder(), []int{10, 9, 13, 8, 0, 80, 16, 32}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %d, want %d", got, want)
	}
}
