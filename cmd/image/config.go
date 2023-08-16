package image

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

type Image struct {
	Registry  string
	Namespace string
	Name      string
	VERSION   string
}

type K8S struct {
	Namespace     string // k8s命名空间
	ResourceType  string // k8s资源类型
	ContainerName string // 容器名称
}
