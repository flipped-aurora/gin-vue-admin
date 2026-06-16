package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

func addManifestFlags(cmd *cobra.Command, params []ManifestParameter) {
	for _, param := range params {
		cmd.Flags().String(param.Flag, "", param.Description)
		if param.Required {
			_ = cmd.MarkFlagRequired(param.Flag)
		}
	}
}

func runManifestCommand(def ManifestCommand, cfg *CliConfig, cmd *cobra.Command) error {
	payload, path, err := buildRequestPayload(def, cmd)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(strings.ToUpper(def.Method), strings.TrimRight(cfg.BaseURL, "/")+path, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	if cfg.Token != "" {
		header := cfg.AuthHeader
		if header == "" {
			header = "x-token"
		}
		req.Header.Set(header, cfg.Token)
	}
	if len(payload) > 0 {
		contentType := def.ContentType
		if contentType == "" {
			contentType = "application/json"
		}
		req.Header.Set("Content-Type", contentType)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed: %s %s", resp.Status, strings.TrimSpace(string(body)))
	}
	_, err = io.Copy(cmd.OutOrStdout(), resp.Body)
	return err
}

func buildRequestPayload(def ManifestCommand, cmd *cobra.Command) ([]byte, string, error) {
	path := def.Path
	queryParts := make([]string, 0)
	body := make(map[string]any)
	for _, param := range def.Parameters {
		value, err := cmd.Flags().GetString(param.Flag)
		if err != nil {
			return nil, "", err
		}
		if strings.TrimSpace(value) == "" {
			continue
		}
		switch param.Location {
		case "path":
			path = strings.ReplaceAll(path, ":"+param.Field, value)
			path = strings.ReplaceAll(path, "{"+param.Field+"}", value)
		case "query":
			queryParts = append(queryParts, fmt.Sprintf("%s=%s", param.Field, value))
		case "body", "formData":
			body[param.Field] = parseFlagValue(param.Type, value)
		case "header":
			// reserved for future explicit headers
		default:
			body[param.Field] = parseFlagValue(param.Type, value)
		}
	}
	if len(queryParts) > 0 {
		sep := "?"
		if strings.Contains(path, "?") {
			sep = "&"
		}
		path += sep + strings.Join(queryParts, "&")
	}
	if len(body) == 0 {
		return nil, path, nil
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, "", err
	}
	return payload, path, nil
}

func parseFlagValue(typeName string, value string) any {
	switch strings.ToLower(typeName) {
	case "integer", "int", "number":
		return json.Number(value)
	case "boolean", "bool":
		return strings.EqualFold(value, "true")
	default:
		return value
	}
}
