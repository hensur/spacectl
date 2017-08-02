package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mittwald/spacectl/view"
	"os"
	"github.com/mittwald/spacectl/cmd/helper"
)

// spacesShowCmd represents the show command
var spacesShowCmd = &cobra.Command{
	Use:   "show -t <team> <space-name>",
	Short: "Show details regarding a specific space",
	Long: "Show details regarding a specific space",
	RunE: func(cmd *cobra.Command, args []string) error {
		space, err := helper.GetSpaceFromContext(args, spaceFile, api)
		if err != nil {
			RootCmd.SilenceUsage = false
			return err
		}

		v := view.TabularSpaceDetailView{}
		v.SpaceDetail(space, os.Stdout)

		return nil
	},
}

func init() {
	spacesCmd.AddCommand(spacesShowCmd)
}