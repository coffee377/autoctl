package types

import (
	"fmt"
	"strings"
)

type StringValue interface {
	ToString() string
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
