package main

import (
	"path/filepath"

	"github.com/tinned-fish/gofish/internal/home"
)

type tank map[string]string

func (t tank) fill() {
	t["GOFISH_HOME"] = home.String()
	t["GOFISH_CACHE"] = home.Cache()
	t["GOFISH_BARREL"] = home.Barrel()
	t["GOFISH_RIGS"] = home.Rigs()
	t["GOFISH_DEFAULT_RIG"] = filepath.Join(home.Rigs(), home.DefaultRig())
}
