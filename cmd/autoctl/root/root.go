package root

import (
	"os"

	"github.com/coffee377/autoctl/internal/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Options struct {
	cwd       string // 当前工作目录
	directory string // 子目录
	config    string // 配置文件名称
	profile   string // 环境名称
	verbose   bool   // 输出详细信息
}

var (
	rooOpts = &Options{}
	rootCmd = &cobra.Command{
		Use:     "autoctl",
		Aliases: []string{"auto"},
		Long:    "Dev opts automation command line tool",
	}
)

// GetRootCommand returns the root cobra command to be executed by main.
func GetRootCommand() *cobra.Command {
	//c.AddCommand(cmd.GetReleaseCommand())
	//c.AddCommand(cmd.GetVersionCommand())
	rootCmd.AddCommand(cmd.GetTokenCommand())

	return rootCmd
}

func init() {
	// 1. 定义全局持久化 Flag
	rootCmd.PersistentFlags().StringVarP(&rooOpts.cwd, "directory", "C", "", "change execution directory")
	rootCmd.PersistentFlags().StringVarP(&rooOpts.directory, "module-path", "m", "", "change execution directory into submodule path")
	rootCmd.PersistentFlags().StringVarP(&rooOpts.config, "file", "f", "", "config file (default is $HOME/auto.yml)")
	rootCmd.PersistentFlags().BoolVarP(&rooOpts.verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&rooOpts.profile, "profile", "P", "", "config profile")

	_ = viper.BindPFlags(rootCmd.PersistentFlags())
}

func Execute() {
	if err := GetRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
