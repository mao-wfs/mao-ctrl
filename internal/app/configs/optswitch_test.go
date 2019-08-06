package configs_test

import (
	"reflect"
	"testing"

	"github.com/mao-wfs/mao-ctrl/internal/app/configs"
	"github.com/mao-wfs/mao-ctrl/internal/pkg/testutil"
)

func TestLoadOptSwitchConfig(t *testing.T) {
	pairs := map[string]string{
		"PG_HOST":  "127.0.0.1",
		"PG_PORT":  "5000",
		"PG_ORDER": "2048,0,20480,4096,8192",
		"FG_HOST":  "127.0.0.1",
		"FG_PORT":  "5001",
	}

	reset := testutil.SetEnvs(t, pairs)
	defer reset()

	conf, err := configs.LoadOptSwitchConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.PG.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.PG.Port, int16(5000); got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if got, want := conf.PG.Order, []int{2048, 0, 20480, 4096, 8192}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %d, want %d", got, want)
	}
	if got, want := conf.FG.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.FG.Port, int16(5001); got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestLoadOptSwitchConfigOrderDefault(t *testing.T) {
	pairs := map[string]string{
		"PG_HOST": "127.0.0.1",
		"PG_PORT": "5000",
		"FG_HOST": "127.0.0.1",
		"FG_PORT": "5001",
	}

	reset := testutil.SetEnvs(t, pairs)
	defer reset()

	conf, err := configs.LoadOptSwitchConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := conf.PG.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.PG.Port, int16(5000); got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if got, want := conf.PG.Order, []int{2500, 2304, 3328, 2048, 0, 20480, 4096, 8192}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %d, want %d", got, want)
	}
	if got, want := conf.FG.Host, "127.0.0.1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := conf.FG.Port, int16(5001); got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestOptSwitchConfig(t *testing.T) {
	conf := &configs.OptSwitchConfig{
		PG: &configs.PGConfig{
			Host:  "127.0.0.1",
			Port:  5000,
			Order: []int{2048, 0, 20480, 4096, 8192},
		},
		FG: &configs.FGConfig{
			Host: "127.0.0.1",
			Port: 5001,
		},
	}

	t.Run("get PG address", func(t *testing.T) {
		if got, want := conf.PGAddr(), "127.0.0.1:5000"; got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	})

	t.Run("get switching order", func(t *testing.T) {
		if got, want := conf.Order(), []int{2048, 0, 20480, 4096, 8192}; !reflect.DeepEqual(got, want) {
			t.Fatalf("got %d, want %d", got, want)
		}
	})

	t.Run("get FG address", func(t *testing.T) {
		if got, want := conf.FGAddr(), "127.0.0.1:5001"; got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	})
}
