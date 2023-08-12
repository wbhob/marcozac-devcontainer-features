package devcontainerfeatures

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/marcozac/devcontainer-features/pkg/feature"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

//go:embed template/README.md.tmpl
var readmeTmpl string

var generateReadmeCmd = &cobra.Command{
	Use:   "generate-readme",
	Short: "Generate the README.md file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		srcDir := cmd.Flag("src-dir").Value.String()
		dirs, err := os.ReadDir(srcDir)
		if err != nil {
			return errors.Wrap(err, "read src directory")
		}

		feats := make([]*feature.Feature, 0, len(dirs))
		for _, dir := range dirs {
			if !dir.IsDir() {
				continue
			}
			feat, err := feature.FromDirPath(filepath.Join("./src", dir.Name()))
			if err != nil {
				return errors.Wrap(err, "read feature")
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
			return errors.Wrap(err, "execute template")
		}

		if cmd.Flag("dry-run").Value.String() == "true" {
			os.Stdout.Write(buf.Bytes())
			return nil
		}

		filePath := cmd.Flag("output").Value.String()
		f, err := os.Create(filePath)
		if err != nil {
			return errors.Wrapf(err, "create %s", filePath)
		}
		defer f.Close()

		if _, err := f.Write(buf.Bytes()); err != nil {
			return errors.Wrapf(err, "write %s", filePath)
		}

		return nil
	},
}

func init() {
	generateReadmeCmd.Flags().String("src-dir", "./src", "The path to the src directory.")
	generateReadmeCmd.Flags().StringP("output", "o", "./README.md", "The README.md file output.")
	generateReadmeCmd.Flags().Bool("dry-run", false, "Print the generated README.md to stdout instead of writing it to disk.")
}
