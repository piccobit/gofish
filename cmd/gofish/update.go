package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/piccobit/gofish/pkg/home"
	"github.com/piccobit/gofish/pkg/ohai"
	"github.com/piccobit/gofish/pkg/rig/installer"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update rigs",
		RunE: func(_ *cobra.Command, _ []string) error {
			return updateRigs()
		},
	}
	return cmd
}

func updateRigs() error {
	start := time.Now()
	rigs := findRigs(home.Rigs())
	for _, rig := range rigs {
		i, err := installer.FindSource(filepath.Join(home.Rigs(), rig))
		if err != nil {
			return err
		}
		if err := installer.Update(i); err != nil {
			return err
		}
	}
	t := time.Now()
	ohai.Ohailn("Rigs updated!")
	table := uitable.New()
	table.AddRow("NAME")
	for _, rig := range rigs {
		table.AddRow(rig)
	}
	fmt.Printf("%s\n\n", table)
	ohai.Successf("rigs updated in %s\n", t.Sub(start).String())
	return nil
}
