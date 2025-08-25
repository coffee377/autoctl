package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetVersionCommand() (versionCmd *cobra.Command) {
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Increments the version number according to the semantic version",
		Run: func(cmd *cobra.Command, args []string) {
			// todo version called
			fmt.Println("version called")
		},
	}
	return versionCmd
}
