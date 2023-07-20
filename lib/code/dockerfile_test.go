package code

import "testing"

func TestImage_ToString(t *testing.T) {

	tag := "harbor.jqk8s.jqsoft.net/zhmz/node:16.20.0-alpine"
	sha256 := "harbor.jqk8s.jqsoft.net/zhmz/node@sha256:710a2c192ca426e03e4f3ec1869e5c29db855eb6969b74e6c50fd270ffccd3f1"

	node1 := Image{
		Registry:  "harbor.jqk8s.jqsoft.net",
		Namespace: "zhmz",
		Name:      "node",
		Tag:       "16.20.0-alpine",
	}
	node2 := Image{
		Registry:  "harbor.jqk8s.jqsoft.net",
		Namespace: "zhmz",
		Name:      "node",
		Tag:       "1.20.1-alpine",
		Digest:    "sha256:710a2c192ca426e03e4f3ec1869e5c29db855eb6969b74e6c50fd270ffccd3f1",
	}

	if tag != node1.ToString() {
		t.Errorf("Expectd: %s, but got %s", tag, node1.ToString())
	}

	if sha256 != node2.ToString() {
		t.Errorf("Expectd: %s, but got %s", sha256, node2.ToString())
	}

}

func TestDockerfile_Render(t *testing.T) {
	nodeImage := Image{
		Registry:  "harbor.jqk8s.jqsoft.net",
		Namespace: "zhmz",
		Name:      "node",
		Tag:       "16.20.0-alpine",
	}
	nginxImage := Image{
		Registry:  "harbor.jqk8s.jqsoft.net",
		Namespace: "zhmz",
		Name:      "node",
		Tag:       "1.20.1-alpine",
	}
	dockerfile := &Dockerfile{
		Args: []Parameter{
			{"NODE_IMAGE", nodeImage, "NodeJS镜像"},
			{"NGINX_IMAGE", nginxImage, "Nginx镜像"},
			{"TEST", nil, ""},
		},
		Stages: []Stage{
			{Name: "builder", Description: "", Image: Image{}, ImageArg: "NODE_IMAGE", Args: []Parameter{}},
			{Name: "web", Description: "", Image: Image{}, ImageArg: "NODE_IMAGE", Args: []Parameter{}},
		},
	}
	dockerfile.Render()
}
