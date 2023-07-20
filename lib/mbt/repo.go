package main

import (
	git "github.com/libgit2/git2go/v34"
)

type gitBlob struct {
	path string
	//commit *gitCommit
	entry *git.TreeEntry
}

func main() {
	//git
	repository, err := git.OpenRepository("D:\\Project\\jqsoft\\framework\\frontend-base")
	if err != nil {
		println(repository)
	}
}
