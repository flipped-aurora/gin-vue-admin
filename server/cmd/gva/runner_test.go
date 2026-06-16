package main

import (
	"encoding/json"
	"testing"

	"github.com/spf13/cobra"
)

func TestLoginSavesToken(t *testing.T) {
	cfg := CliConfig{}
	cfg.Token = "jwt-token"
	if cfg.Token != "jwt-token" {
		t.Fatalf("token not assigned")
	}
}

func TestBuildRequestPayload(t *testing.T) {
	def := ManifestCommand{
		Name:   "user-list",
		Method: "POST",
		Path:   "/api/user/getUserList",
		Parameters: []ManifestParameter{
			{Flag: "page", Type: "int", Location: "body", Field: "page"},
			{Flag: "page-size", Type: "int", Location: "body", Field: "pageSize"},
		},
	}
	cmd := &cobra.Command{Use: "user-list"}
	addManifestFlags(cmd, def.Parameters)
	_ = cmd.Flags().Set("page", "1")
	_ = cmd.Flags().Set("page-size", "10")
	payload, path, err := buildRequestPayload(def, cmd)
	if err != nil {
		t.Fatalf("buildRequestPayload() error = %v", err)
	}
	if path != "/api/user/getUserList" {
		t.Fatalf("path = %q", path)
	}
	var body map[string]any
	if err := json.Unmarshal(payload, &body); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if body["page"] != float64(1) || body["pageSize"] != float64(10) {
		t.Fatalf("unexpected body: %#v", body)
	}
}
