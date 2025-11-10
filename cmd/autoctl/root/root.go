package root

import (
	"fmt"
	"os"
	"path"

	"github.com/coffee377/autoctl/internal/cmd"
	"github.com/coffee377/autoctl/pkg/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetRootCommand returns the root cobra command to be executed by main.
func GetRootCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "autoctl",
		Aliases: []string{"auto"},
		Long:    "Dev opts automation command line tool",
	}

	c.AddCommand(cmd.GetReleaseCommand())
	c.AddCommand(cmd.GetVersionCommand())

	return c
}

type Options struct {
	cwd       string // 当前工作目录
	directory string // 子目录
	config    string // 配置文件名称
	verbose   bool   // 输出详细信息
}

var (
	rootCmd = GetRootCommand()
	rooOpts = &Options{}
)

func init() {
	cobra.OnInitialize(loadConfig)
	rootCmd.PersistentFlags().StringVarP(&rooOpts.config, "file", "f", "",
		"config file (default is $HOME/auto.yml)")
	rootCmd.PersistentFlags().StringVarP(&rooOpts.cwd, "directory", "C", "",
		"change execution directory")
	rootCmd.PersistentFlags().StringVarP(&rooOpts.directory, "--module-path", "m", "",
		"change execution directory into submodule path")
	rootCmd.PersistentFlags().BoolVarP(&rooOpts.verbose, "verbose", "v", false,
		"verbose output")

	//image.RegisterCommandRecursive(rootCmd, image.RootOptions{})
	//version.RegisterCommandRecursive(rootCmd)
}

func loadConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if rooOpts.config != "" {
		// Use config file from the flag.
		viper.SetConfigFile(rooOpts.config)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "server" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")                    // 当前执行目录寻找配置文件
		viper.AddConfigPath(path.Join(".", "conf")) // 当前执行目录 conf 目录下寻找配置文件
		viper.SetConfigName("auto")
		viper.AutomaticEnv() // read in environment variables that match
	}

	_ = viper.ReadInConfig()
	configFile := viper.ConfigFileUsed()

	if configFile != "" && log.IsDebugEnabled() {
		log.Debug("Using config file: %s", configFile)
	}

	//_ = viper.Unmarshal(&serverOption)

}

func Execute() {
	if err := GetRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
