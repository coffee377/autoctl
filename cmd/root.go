package cmd

import (
	"bytes"
	"fmt"
	"github.com/coffee377/autoctl/cmd/image"
	"github.com/coffee377/autoctl/cmd/version"
	"github.com/coffee377/autoctl/pkg/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"text/template"
)

type RootOption struct {
	ConfigFile string `mapstructure:"configFile"` // 配置文件目录
	Cwd        string `mapstructure:"cwd"`
}

const (
	helpTemplate = `{{insertTemplate . (or .Long .Short) | trimTrailingWhitespaces}}

{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`
	usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{insertTemplate . .Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
)

// EnableUsageTemplating enables gotemplates for usage strings, i.e. cmd.Short, cmd.Long, and cmd.Example.
// The data for the template is the command itself. Especially useful are `.Root.Name` and `.CommandPath`.
// This will be inherited by all subcommands, so enabling it on the root command is sufficient.
func EnableUsageTemplating(cmds ...*cobra.Command) {
	cobra.AddTemplateFunc("insertTemplate", TemplateCommandField)
	for _, cmd := range cmds {
		cmd.SetHelpTemplate(helpTemplate)
		cmd.SetUsageTemplate(usageTemplate)
	}
}

var usageTemplateFuncs = template.FuncMap{}

// AddUsageTemplateFunc adds a template function to the usage template.
func AddUsageTemplateFunc(name string, f interface{}) {
	usageTemplateFuncs[name] = f
}

func TemplateCommandField(cmd *cobra.Command, field string) (string, error) {
	t := template.New("")
	t.Funcs(usageTemplateFuncs)
	t, err := t.Parse(field)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	if err := t.Execute(&out, cmd); err != nil {
		return "", err
	}
	return out.String(), nil
}

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "autoctl",
		Aliases: []string{"auto"},
		Short:   "Dev opts automation command line tool",
		//TraverseChildren: true,
		//		Example: `
		//autoctl version -C packages/teamwork-ui
		//autoctl version -C packages/teamwork-ui
		//autoctl version -C packages/teamwork-ui
		//`,
	}
	os.Getenv("log.level")

	//EnableUsageTemplating(cmd)

	return cmd
}

var (
	rootCmd = NewRootCmd()
	rooOpts = &RootOptions{}
)

type RootOptions struct {
	cwd       string // 当前工作目录
	directory string // 子目录
	config    string // 配置文件名称
	verbose   bool   // 输出详细信息
}

func init() {
	cobra.OnInitialize(loadConfig)
	rootCmd.PersistentFlags().StringVarP(&rooOpts.config, "file", "f", "", "config file (default is $HOME/auto.yml)")
	rootCmd.PersistentFlags().StringVarP(&rooOpts.cwd, "directory", "C", "", "change execution directory")
	rootCmd.PersistentFlags().StringVarP(&rooOpts.directory, "--module-path", "m", "", "change execution directory into submodule path")
	rootCmd.PersistentFlags().BoolVarP(&rooOpts.verbose, "verbose", "v", false, "verbose output")

	image.RegisterCommandRecursive(rootCmd, image.RootOptions{})
	version.RegisterCommandRecursive(rootCmd)
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
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
