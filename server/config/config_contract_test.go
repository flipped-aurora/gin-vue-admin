package config

import (
	"bytes"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestShippedServerConfigsMatchSchema(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		kubernetes bool
	}{
		{name: "default", path: "../config.yaml"},
		{name: "docker", path: "../config.docker.yaml"},
		{name: "kubernetes", path: "../../deploy/kubernetes/server/gva-server-configmap.yaml", kubernetes: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := os.ReadFile(tt.path)
			if err != nil {
				t.Fatal(err)
			}
			if tt.kubernetes {
				var manifest struct {
					Data map[string]string `yaml:"data"`
				}
				if err := yaml.Unmarshal(content, &manifest); err != nil {
					t.Fatal(err)
				}
				embedded, ok := manifest.Data["config.yaml"]
				if !ok || len(bytes.TrimSpace([]byte(embedded))) == 0 {
					t.Fatal("Kubernetes ConfigMap 缺少非空 data.config.yaml")
				}
				content = []byte(embedded)
			}

			decoder := yaml.NewDecoder(bytes.NewReader(content))
			decoder.KnownFields(true)
			var cfg Server
			if err := decoder.Decode(&cfg); err != nil {
				t.Fatalf("配置模板与 config.Server 不一致: %v", err)
			}
		})
	}
}
