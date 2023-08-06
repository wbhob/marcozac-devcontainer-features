package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/marcozac/devcontainer-features/pkg/feature"

	_ "embed"
)

//go:embed template/README.md.tmpl
var readmeTmpl string

func main() {
	var dryRun bool
	flag.BoolVar(&dryRun, "dry-run", false, "Print the generated README.md to stdout instead of writing it to disk.")
	flag.Parse()

	dirs, err := os.ReadDir("./src")
	if err != nil {
		log.Fatal("generate readme: read src directory: ", err)
	}
	feats := make([]*feature.Feature, 0, len(dirs))
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		feat, err := feature.FromDirPath(filepath.Join("./src", dir.Name()))
		if err != nil {
			log.Fatalf("generate readme: %s: %s", dir, err)
		}
		feats = append(feats, feat)
	}
	var buf bytes.Buffer
	err = template.Must(template.New("README.md.tmpl").
		Funcs(template.FuncMap{
			"sub": func(a, b int) int { return a - b },
		}).
		Parse(readmeTmpl),
	).Execute(&buf, feats)
	if err != nil {
		log.Fatal("generate readme: ", err)
	}
	if dryRun {
		os.Stdout.Write(buf.Bytes())
		return
	}
	f, err := os.Create("./README.md")
	if err != nil {
		log.Fatal("generate readme: create README.md: ", err)
	}
	defer f.Close()
	if _, err := f.Write(buf.Bytes()); err != nil {
		f.Close()
		log.Fatal("generate readme: write README.md: ", err)
	}
}
