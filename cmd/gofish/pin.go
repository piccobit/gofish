package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tinned-fish/gofish/internal/home"
	"github.com/tinned-fish/gofish/internal/ohai"
)

const (
	DotPinned = ".pinned"
)

func newPinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pin <food>",
		Short: "protect a fish food, preventing fish from installing upgrades",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if Pinned(args[0]) {
				ohai.Ohaif("%s is already pinned\n", args[0])

				return nil
			} else {
				if err := ioutil.WriteFile(
					filepath.Join(home.Barrel(), args[0], DotPinned), []byte{}, 0644); err != nil {
					return err
				}

				ohai.Ohaif("%s is pinned\n", args[0])
			}

			return nil
		},
	}
	return cmd
}

func Pinned(name string) bool {
	files, err := ioutil.ReadDir(filepath.Join(home.Barrel(), name))
	if err != nil {
		return false
	}

	for _, f := range files {
		if !f.IsDir() && (f.Name() == DotPinned) {
			return true
		}
	}

	return false
}
