package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tinned-fish/gofish/internal/home"
	"github.com/tinned-fish/gofish/internal/ohai"
)

func newUnpinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpin <food>",
		Short: "remove protection from a fish food, allowing fish to install upgrades",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !Pinned(args[0]) {
				ohai.Ohaif("%s is already unpinned\n", args[0])

				return nil
			} else {
				if err := os.Remove(filepath.Join(home.Barrel(), args[0], DotPinned)); err != nil {
					return err
				}

				ohai.Ohaif("%s unpinned\n", args[0])
			}

			return nil
		},
	}
	return cmd
}
