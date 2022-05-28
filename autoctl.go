package main

import (
	"github.com/coffee377/autoctl/cmd"
	"github.com/coffee377/autoctl/dingtalk"
)

func main() {
	cmd.Execute()
	dingtalk.GetAccessToken()
}
