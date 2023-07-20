package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	git "github.com/coffee377/autoctl/git/commit"
	"github.com/coffee377/autoctl/log"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

type Plus struct {
	Cwd     string
	Verbose bool
	name    string // 模块名称
	path    string // 模块路径
	version string // 版本
	args    *[]string
	cmd     *cobra.Command
}

// Exec 执行 git 命令
func (plus *Plus) Exec(args ...string) []byte {
	command := exec.Command("git", args...)
	command.Dir = plus.Cwd
	if plus.Verbose {
		n := strings.SplitN(command.String(), " ", 2)
		log.Debug("git %s", n[1])
	}
	output, err := command.Output()
	if err != nil {
		log.Fatal("%s", err)
	}
	return output
}

// FetchAll 拉取所有远程仓库的最新内容到本地
func (plus *Plus) FetchAll() {
	_ = plus.Exec("fetch", "--all")
}

// FetchLatestTag 获取最近一次的标签
func (plus *Plus) FetchLatestTag() string {
	return string(plus.Exec("describe", "--tags", "--abbrev=0"))
}

// https://zhuanlan.zhihu.com/p/87725726 打标签
// https://blog.csdn.net/qq_21746331/article/details/120776710

type Tag struct {
	objecttype  string // The type of the object (blob, tree, commit, tag)
	objectname  string // The object name (aka SHA-1)
	creatordate string //
	refname     string // short tag name
	subject     string // tag des
}

func (plus *Plus) FetchTags(desc bool, patterns ...string) []byte {
	var args []string
	var min = ""
	if desc {
		min = "-"
	}

	var sort = fmt.Sprintf("--sort=%screatordate", min)

	args = append(args, "tag", "-l", sort, "--format=%(objecttype);%(objectname);%(creatordate:short);%(refname:short);%(subject)")
	// pattern 参数需要放在最后
	for _, value := range patterns {
		args = append(args, value)
	}
	// git tag -l --sort=-creatordate --format='%(objecttype);%(objectname);%(creatordate:short);%(refname:short);%(subject)'
	bytes := plus.Exec(args...)
	s := string(bytes)
	log.Info(s)

	return nil
}

// FetchLogs will query commit records between two tags
func (plus *Plus) FetchLogs(commit1, commit2 string) []byte {
	//var notation string
	//if len(tag1) > 0 && len(tag2) > 0 {
	//	notation = fmt.Sprintf("%s..%s", tag1, tag2)
	//} else if len(tag1) > 0 {
	//	notation = tag1
	//} else if len(tag2) > 0 {
	//	notation = tag2
	//}
	var (
		args    []string
		records []*git.CommitRecord
	)

	format := fmt.Sprintf("--pretty=%s", git.RecordFormat)
	args = append(args, "log", "--no-merges", "--date=format:%Y/%m/%d %H:%M:%S", format)

	args = append(args, "-i")
	args = append(args, "-E", "--grep", "fix|zap")

	result := plus.Exec(args...)
	logs := bytes.Split(result, []byte(git.LogSep))

	for _, v := range logs {
		if len(v) == 0 {
			continue
		}
		record := git.NewCommitRecord(v)
		records = append(records, record)
	}

	marshal, _ := json.Marshal(records)
	// 追加换行符
	marshal = append(marshal, []byte("\n")...)
	return marshal
}
