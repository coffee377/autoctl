package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCmd() (versionCmd *cobra.Command) {
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Increments the version number according to the semantic version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version called")
		},
	}
	return versionCmd
}

func RegisterCommandRecursive(parent *cobra.Command) {
	versionCmd := NewVersionCmd()
	parent.AddCommand(versionCmd)
}
