package main

import (
	"time"

	"github.com/piccobit/gofish/pkg/ohai"
	"github.com/piccobit/gofish/pkg/rig/installer"
	"github.com/spf13/cobra"
)

func newRigAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <rig> [name]",
		Short: "add rigs",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := ""
			if len(args) > 1 {
				name = args[1]
			}
			i, err := installer.New(args[0], name, "")
			if err != nil {
				return err
			}

			start := time.Now()
			if err := installer.Install(i); err != nil {
				return err
			}
			t := time.Now()
			ohai.Successf("rig constructed in %s\n", t.Sub(start).String())
			return nil
		},
	}
	return cmd
}
