package code

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type StringValue interface {
	ToString() string
}

type Parameter struct {
	Name        string
	Value       interface{}
	Description string
}

type Image struct {
	Registry  string `json:"registry"`  // image registry
	Namespace string `json:"namespace"` // image namespace
	Name      string `json:"name"`      // image name
	Tag       string `json:"tag"`       // image tag
	Digest    string `json:"digest"`    // image digest in the way sha256:aa.... Please note this parameter, if set, will override the tag
}

func (i Image) ToString() string {
	elems := make([]string, 0)

	if i.Registry != "" {
		elems = append(elems, i.Registry)
	}

	repository := fmt.Sprintf("%s/%s", i.Namespace, i.Name)
	if repository != "" {
		elems = append(elems, repository)
	}

	tag := fmt.Sprintf(":%s", i.Tag)
	if i.Digest != "" {
		tag = fmt.Sprintf("@%s", i.Digest)
	}

	return strings.Join(elems, "/") + tag
}

type Stage struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       Image       `json:"image"`
	ImageArg    string      `json:"image_arg"`
	Args        []Parameter `json:"args"`
}

type Dockerfile struct {
	Args   []Parameter `json:"args"`
	Stages []Stage     `json:"stages"`
}

func (dockerfile *Dockerfile) Render() {
	// 创建一个模板函数映射
	funcMap := template.FuncMap{
		"ValidArgs": ValidArgs,
	}
	// 创建模板并注册函数映射
	tmpl := template.New("dockerfile.tmpl").Funcs(funcMap)
	// 解析模板内容
	t, err := tmpl.ParseFiles("./dockerfile.tmpl")
	if err != nil {
		fmt.Println("Pare template failed err =", err)
		return
	}
	err = t.Execute(os.Stdout, dockerfile)
	if err != nil {
		fmt.Println("render template failed err =", err)
		return
	}
}

func ValidArgs(args []Parameter) []Parameter {
	res := make([]Parameter, 0)
	for _, arg := range args {
		if arg.Name != "" && arg.Value != nil {
			res = append(res, arg)
		}
	}
	return res
}
