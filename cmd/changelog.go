package cmd

import (
	"github.com/coffee377/autoctl/git"
	"github.com/spf13/cobra"
)

const (
	JSON     = "json"
	MARKDOWN = "markdown"
	HTML     = "html"
)

type changeLogOptions struct {
	fetch  bool
	format string
	output bool
}

var logOpts = changeLogOptions{
	fetch:  false,
	format: "json",
	output: false,
}

// changelogCmd represents the changelog command
var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "command for git to generate logs",
	Run: func(cmd *cobra.Command, args []string) {
		//var err error
		gitPlus := git.Plus{Cwd: rootOpts.cwd, Verbose: rootOpts.verbose}
		// 1. 拉取最新代码
		if logOpts.fetch {
			gitPlus.FetchAll()
		}
		// 获取 log
		c, e := gitPlus.FetchLogs("", "")
		_, _ = cmd.OutOrStdout().Write(c)
		_, _ = cmd.OutOrStderr().Write(e)
	},
}

func init() {
	rootCmd.AddCommand(changelogCmd)

	// 是否获取所有远程参考最新代码
	changelogCmd.Flags().BoolVar(&logOpts.fetch, "fetch", false, "fetch the latest commits and tags for all branches")
	// 输出格式 json markdown html
	changelogCmd.Flags().StringVarP(&logOpts.format, "format", "f", "json", "output file format json, markdown or html")
	// 是否输出到文件，默认 false ,输出到控制台
	changelogCmd.Flags().BoolVarP(&logOpts.output, "output", "o", false, "whether to output logs as file")
}

//type ChangeLog struct {
//	Tag     string         `json:"tag"`     // tag 标签, 最近的标签到最新的提交为 Unreleased
//	Commits []CommitRecord `json:"commits"` //
//}
