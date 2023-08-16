package release

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewReleaseCmd() (releaseCmd *cobra.Command) {
	releaseCmd = &cobra.Command{
		Use:   "release",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("release called")
		},
	}
	return releaseCmd
}

func RegisterCommandRecursive(parent *cobra.Command) {
	versionCmd := NewReleaseCmd()
	parent.AddCommand(versionCmd)
}
