package image

type Image struct {
	Registry  string `json:"registry"`  // image registry
	Namespace string `json:"namespace"` // image namespace
	Name      string `json:"name"`      // image name
	Tag       string `json:"tag"`       // image tag
	Digest    string `json:"digest"`    // image digest in the way sha256:aa.... Please note this parameter, if set, will override the tag
}

type K8S struct {
	Namespace     string // k8s命名空间
	ResourceType  string // k8s资源类型
	ContainerName string // 容器名称
}

type Conf struct {
	Global Image
	//Name  string
	//Image Image
	//K8S   K8S
	cc map[string]sss
}

type sss struct {
	Image Image
	K8S   K8S
}
