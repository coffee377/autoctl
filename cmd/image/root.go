package image

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/coffee377/autoctl/pkg/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RootOptions struct {
	cwd     string // 当前工作目录
	config  string // 配置文件名称
	verbose bool   // 输出详细信息
	name    string // 配置项
}

func NewImageCmd(opts RootOptions) (imageCmd *cobra.Command) {
	imageCmd = &cobra.Command{
		Use:   "image",
		Short: "Run the image tools",
		//Run:   func(cmd *cobra.Command, args []string) {},
		RunE: func(cmd *cobra.Command, args []string) error {
			//ctx := cmd.Context()
			return nil
		},
		TraverseChildren: true,
	}

	initConfig(opts.config, "image")

	imageCmd.PersistentFlags().StringVarP(&opts.cwd, "cwd", "p", "", "set the current working directory")
	imageCmd.PersistentFlags().StringVarP(&opts.config, "config", "c", "", "config file (default is $HOME/image.yaml)")
	imageCmd.PersistentFlags().BoolVarP(&opts.verbose, "verbose", "v", false, "verbose output")

	imageCmd.AddCommand(NewImagePullCmd())

	return imageCmd
}

func RegisterCommandRecursive(parent *cobra.Command, opts RootOptions) {
	parent.AddCommand(NewImageCmd(opts))
}

func initConfig(cfgFile, configName string) {
	var (
		home      = ""
		binDir, _ = filepath.Split(os.Args[0])
		err       error
	)

	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err = homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(binDir)
		viper.AddConfigPath(path.Join(binDir, "conf"))
		viper.AddConfigPath(path.Join(binDir, "../conf"))
		viper.SetConfigName(configName)
	}

	if err = viper.ReadInConfig(); err != nil {
		//fmt.Println("Can't read config:", err)
		log.Info("初始化默认配置")

		//filename := path.Join(home, "automation.yaml")
		//err = viper.WriteConfigAs(filename)
		//if err != nil {
		//	return
		//}
		//CliOpts.Verbose = true
		//err := viper.WriteConfig()
		//if err != nil {
		//	return
		//}
		//err := ioutil.WriteFile(filename, nil, 0666)
		//if err != nil {
		//	os.Exit(1)
		//}

	}
}
