package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "autoctrl",
	Short:            "autoctrl command tool",
	Run:              func(cmd *cobra.Command, args []string) {},
	TraverseChildren: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type rootOptions struct {
	cwd        string // 当前工作目录
	configFile string // 配置文件名称
	verbose    bool   // 输出详细信息
	//changelog  ChangeLogOptions
}

var rootOpts = rootOptions{
	//configFile: "",
	//verbose:    false,
	//changelog:  ChangeLogOptions{},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&rootOpts.cwd, "cwd", "p", "", "set the current working directory")
	rootCmd.PersistentFlags().StringVarP(&rootOpts.configFile, "config", "c", "", "config file (default is $HOME/automation.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&rootOpts.verbose, "verbose", "v", false, "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	var (
		home       = ""
		binDir, _  = filepath.Split(os.Args[0])
		configName = "automation"
		err        error
	)

	// Don't forget to read config either from cfgFile or from home directory!
	cfgFile := rootOpts.configFile
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

		// Search config in home directory with name "automation" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(binDir)
		viper.AddConfigPath(path.Join(binDir, "conf"))
		viper.SetConfigName(configName)
	}

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		fmt.Println("初始化默认配置")

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
