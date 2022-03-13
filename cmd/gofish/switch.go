package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func newSwitchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "switch <food> <version>",
		Short: "switch fish food to another version",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := getFood(args[0])
			if err != nil {
				return err
			}

			// oldVersion := f.Version

			pkg := f.GetPackage(runtime.GOOS, runtime.GOARCH)
			if pkg == nil {
				return fmt.Errorf("food '%s' does not support the current platform (%s/%s)", f.Name, runtime.GOOS, runtime.GOARCH)
			}

			if err := f.Unlink(pkg); err != nil {
				return err
			}

			versions := findFoodVersions(args[0])
			if len(versions) >= 1 {
				for _, version := range versions {
					if version == args[1] {
						f.Version = version

						return f.Link(pkg)
					}
				}

				return fmt.Errorf("version '%s' of food '%s' not available", args[1], args[0])
			}

			return nil
		},
	}
	return cmd
}
