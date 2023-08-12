package devcontainerfeatures

import (
	"os"
	"path"
	"path/filepath"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"

	"github.com/marcozac/devcontainer-features/pkg/feature"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new id",
	Short: "Generate a new feature.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		feat := feature.Feature{
			ID:          id,
			Version:     "0.0.0",
			Name:        cmd.Flag("name").Value.String(),
			Description: cmd.Flag("description").Value.String(),
		}
		if repoSrcURL := cmd.Flag("repo-src-url").Value.String(); repoSrcURL != "" {
			feat.DocumentationURL = path.Join(repoSrcURL, id)
		}

		srcDir := cmd.Flag("src-dir").Value.String()
		featDir := filepath.Join(srcDir, id)
		if err := os.MkdirAll(featDir, 0o755); err != nil {
			return errors.Wrapf(err, "create %s directory", id)
		}

		featJSONPath := filepath.Join(featDir, feature.FeatureFileName)
		if _, err := os.Stat(featJSONPath); err == nil {
			return errors.Wrapf(err, "%s already exists", featJSONPath)
		}

		data, err := json.MarshalIndent(feat, "", "    ")
		if err != nil {
			return errors.Wrapf(err, "marshal %s %s", id, feature.FeatureFileName)
		}

		f, err := os.OpenFile(featJSONPath, os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			return errors.Wrapf(err, "create %s", featJSONPath)
		}

		if _, err := f.Write(data); err != nil {
			return errors.Wrapf(err, "write %s", featJSONPath)
		}

		return nil
	},
}

func init() {
	newCmd.Flags().StringP("name", "n", "", "Required. The name of the feature.")
	_ = newCmd.MarkFlagRequired("name")

	newCmd.Flags().StringP("description", "d", "", "The description of the feature.")
	newCmd.Flags().String("src-dir", "./src", "The path to the src directory.")
	newCmd.Flags().String("repo-src-url", os.Getenv("REPO_SRC_URL"), "The repository src directory URL. Uses 'REPO_SRC_URL' env var as default value.\n\tE.g. https://github.com/marcozac/devcontainer-features/tree/main/src\n")
}
