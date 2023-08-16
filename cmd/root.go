package cmd

import (
	"errors"
	"fmt"
	"github.com/coffee377/autoctl/cmd/image"
	"github.com/coffee377/autoctl/cmd/version"
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
	"os"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "autoctl",
	}

	cmdx.EnableUsageTemplating(cmd)

	image.RegisterCommandRecursive(cmd, image.RootOptions{})
	version.RegisterCommandRecursive(cmd)

	return cmd
}

func Execute() {
	c := NewRootCmd()

	if err := c.Execute(); err != nil {
		if !errors.Is(err, cmdx.ErrNoPrintButFail) {
			_, _ = fmt.Fprintln(c.ErrOrStderr(), err)
		}
		os.Exit(1)
	}
}
