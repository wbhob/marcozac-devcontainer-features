package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/goccy/go-json"

	"github.com/marcozac/devcontainer-features/pkg/feature"
)

const (
	srcPath    = "./src"
	repoSrcURL = "https://github.com/marcozac/devcontainer-features/tree/main/src"
)

func main() {
	var id string
	var name string
	var description string

	flag.StringVar(&id, "id", "", "The id of the feature.")
	flag.StringVar(&name, "name", "", "The name of the feature.")
	flag.StringVar(&description, "description", "", "The description of the feature.")
	flag.Parse()

	switch "" {
	case id:
		fatal("The feature id is required.")
	case name:
		fatal("The feature name is required.")
	}

	feat := feature.Feature{
		ID:               id,
		Version:          "0.0.0",
		Name:             name,
		Description:      description,
		DocumentationURL: path.Join(repoSrcURL, id),
	}

	featDir := filepath.Join(srcPath, id)
	if err := os.MkdirAll(featDir, 0o755); err != nil {
		fatal(fmt.Sprintf("create %s directory: %v", id, err))
	}

	data, err := json.MarshalIndent(feat, "", "    ")
	if err != nil {
		fatal(fmt.Sprintf("marshal %s %s: %v", id, feature.FeatureFileName, err))
	}

	featJSONPath := filepath.Join(featDir, feature.FeatureFileName)
	if _, err := os.Stat(featJSONPath); err == nil {
		fatal(featJSONPath, "already exists")
	}

	f, err := os.OpenFile(featJSONPath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fatal(fmt.Sprintf("create %s %s: %v", id, feature.FeatureFileName, err))
	}

	if _, err := f.Write(data); err != nil {
		fatal(fmt.Sprintf("writing %s %s: %v", id, feature.FeatureFileName, err))
	}
}

func fatal(v ...any) {
	fmt.Println(v...)
	os.Exit(1)
}
