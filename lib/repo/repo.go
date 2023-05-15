package lib

import git "github.com/libgit2/git2go/v34"

type gitBlob struct {
	path string
	//commit *gitCommit
	entry *git.TreeEntry
}
