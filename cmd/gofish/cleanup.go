package main

import (
	"github.com/spf13/cobra"
	"github.com/tinned-fish/gofish/internal/gofish"
)

func newCleanupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cleanup <food...>",
		Short: "cleanup unlinked fish food",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var fudz []*gofish.Food
			for _, arg := range args {
				f, err := getFood(arg)
				if err != nil {
					return err
				}
				fudz = append(fudz, f)
			}
			for _, f := range fudz {
				versions := findFoodVersions(f.Name)
				if len(versions) > 1 {
					for _, ver := range versions {
						f, err := getFood(f.Name)
						if err != nil {
							return err
						}
						f.Version = ver
						if !f.Linked() {
							if err := f.Uninstall(); err != nil {
								return err
							}
						}
					}
				}
			}
			return nil
		},
	}
	return cmd
}
