package cmd

import (
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:     "applications",
	Aliases: []string{"apps"},
	Short:   "View available managed applications",
	Long:    `This command allows you to view applications that are offered as a managed service`,
}

func init() {
	RootCmd.AddCommand(appsCmd)
}
