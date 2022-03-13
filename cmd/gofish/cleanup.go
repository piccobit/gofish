package main

import (
	"github.com/spf13/cobra"
	"github.com/tinned-fish/gofish/internal/gofish"
	"github.com/tinned-fish/gofish/internal/ohai"
)

func newCleanupCmd() *cobra.Command {
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "cleanup <food...>",
		Short: "cleanup unlinked fish food",
		// Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var fudz []*gofish.Food

			if len(args) >= 1 {
				for _, arg := range args {
					f, err := getFood(arg)
					if err != nil {
						return err
					}
					fudz = append(fudz, f)
				}

				for _, f := range fudz {
					err := unlinkVersions(f.Name, dryRun)
					if err != nil {
						return err
					}
				}
			} else {
				for _, f := range findFood() {
					err := unlinkVersions(f, dryRun)
					if err != nil {
						return err
					}
				}
			}

			return nil
		},
	}

	d := cmd.Flags()
	d.BoolVarP(&dryRun, "dry-run", "n", false, "don't cleanup, just show what would be done")

	return cmd
}

func unlinkVersions(n string, dryRun bool) error {
	if Pinned(n) {
		ohai.Ohaif("%s is pinned. Please use `gofish unpin %s` to allow cleanup.\n", n, n)

		return nil
	}

	versions := findFoodVersions(n)
	if len(versions) > 1 {
		for _, ver := range versions {
			f, err := getFood(n)
			if err != nil {
				return err
			}

			f.Version = ver

			if !f.Linked() {
				if dryRun {
					ohai.Ohaif("Would uninstall version '%s' of package '%s'\n", f.Version, f.Name)
				} else {
					if err := f.Uninstall(); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
