package image

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewImagePullCmd() (pullCmd *cobra.Command) {
	pullCmd = &cobra.Command{
		Use: "pull",
		//Short: "Increments the version number according to the semantic version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("image pull")
		},
	}
	return pullCmd
}

type d struct {
}

func pull() {

}
