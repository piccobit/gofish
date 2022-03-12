package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tinned-fish/gofish/version"
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "display version information",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(version.String())
		},
	}
	return cmd
}
