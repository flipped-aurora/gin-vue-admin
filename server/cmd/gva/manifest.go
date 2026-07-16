package main

import (
	"encoding/json"
	"os"
)

func loadManifest(path string) (Manifest, error) {
	manifest := Manifest{}
	data, err := os.ReadFile(path)
	if err != nil {
		return manifest, err
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}
