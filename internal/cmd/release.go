package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GetReleaseCommand() (releaseCmd *cobra.Command) {
	releaseCmd = &cobra.Command{
		Use:   "release",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			// todo release called
			fmt.Println("release called")
		},
	}
	return releaseCmd
}
