package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

const (
	JSON     = "json"
	MARKDOWN = "markdown"
	HTML     = "html"
)

type ChangeLogOptions struct {
	fetch  bool
	format string
	output bool
}

var logOpts = ChangeLogOptions{
	fetch:  false,
	format: "json",
	output: false,
}

// changelogCmd represents the changelog command
var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "command for git to generate logs",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 拉取最新代码
		if logOpts.fetch {
			err := fetchGitRepository()
			if err != nil {
				return
			}
		}
		// 获取 log
		logs, err := getGitLogs("", "")
		if err != nil {
			return
		}
		fmt.Println(logs)
		//fmt.Println(CliOpts)
		//fmt.Println(logOpts)
	},
}

func init() {
	rootCmd.AddCommand(changelogCmd)

	// 是否获取所有分支最新代码
	changelogCmd.Flags().BoolVar(&logOpts.fetch, "fetch", false, "fetch the latest commits and tags for all branches")
	// 输出格式 json markdown html
	changelogCmd.Flags().StringVarP(&logOpts.format, "format", "f", "json", "output file format json, markdown or html")
	// 是否输出到文件，默认 false ,输出到控制台
	changelogCmd.Flags().BoolVarP(&logOpts.output, "output", "o", false, "whether to output logs as file")

}

// 执行 git 命令
func git(cwd string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = cwd
	if CliOpts.Verbose == true {
		fmt.Println(cmd.String())
	}
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(stderr.String())
	}
	return stdout.String(), nil
}

// fetch the latest commits and tags
func fetchGitRepository() error {
	_, err := git(CliOpts.Cwd, "fetch", "--all")
	if err != nil {
		return err
	}
	return nil
}

//
//// getGitTags will query all git tags with created time and subject
//func getGitTags() (string, error) {
//	tags, err := git(*source, "tag", "-n", "-l", "--sort=creatordate", "--format", "%(refname:short);%(creatordate:short);%(subject)")
//	if err != nil {
//		return "", err
//	}
//	return tags, nil
//}
//

type ChangeLog struct {
	Tag     string          `json:"tag"` // tag 标签, 最近的标签到最新的提交为 unrelaseted
	Commits []CommitMessage `json:"commits"`
}

// CommitMessage contains the commit information
type CommitMessage struct {
	CommitHash       string `json:"hash"`       // %H: commit hash
	AbbrevCommitHash string `json:"abbrevHash"` // %h: abbreviated commit hash
	Author           string `json:"author"`     // %an: author name
	Email            string `json:"email"`      // %ae: author email
	Timestamp        int8   `json:"timestamp"`  // %at: author date, UNIX timestamp
	Date             string `json:"date"`       // %ad: author date (format respects --date= option)
	Message          string `json:"message"`    //
	Subject          string
	Body             string
	Footer           string
	Notes            string
}

// getGitLogs will query commit records between two tags
func getGitLogs(tag1, tag2 string) (string, error) {
	//var notation string
	//if len(tag1) > 0 && len(tag2) > 0 {
	//	notation = fmt.Sprintf("%s..%s", tag1, tag2)
	//} else if len(tag1) > 0 {
	//	notation = tag1
	//} else if len(tag2) > 0 {
	//	notation = tag2
	//}
	//format := "--pretty=format:{\"hash\":\"%H\",\"abbrevHash\":\"%h\",\"author\":\"%an\",\"email\":\"%ae\",\"date\":\"%ad\",\"timestamp\":%at,\"message\":\"%f\"}"
	format := "--pretty=format:{\"timestamp\":%at,\"message\":\"%f\"}"
	commits, err := git(CliOpts.Cwd, "log", "--no-merges", "--date=format:%Y-%m-%d %H:%M:%S", format)
	if err != nil {
		return "", err
	}
	return commits, nil
}
