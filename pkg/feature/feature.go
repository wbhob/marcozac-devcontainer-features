package feature

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Feature struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (f *Feature) MajorVersion() string {
	return f.Version[:1]
}

func FromDirPath(dir string) (*Feature, error) {
	data, err := os.ReadFile(filepath.Join(dir, "devcontainer-feature.json"))
	if err != nil {
		return nil, fmt.Errorf("read feature file: %w", err)
	}
	feat := new(Feature)
	if err := json.Unmarshal(data, feat); err != nil {
		return nil, fmt.Errorf("unmarshal feature file: %w", err)
	}
	return feat, nil
}
