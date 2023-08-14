package git

import "github.com/coffee377/autoctl/pkg/git/commit"

type ChangeLog struct {
	NameTags string
	Commits  []git.CommitRecord
}

// Report full report structure
type Report struct {
	pattern   string
	originURL string
	commitON  bool
	authorON  bool
	ChangeLog []ChangeLog
	JSON      string
	Markdown  string
	HTML      string
}
