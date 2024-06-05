package cmd

import (
	"github.com/roberthamel/git-repos/internal/logger"
	"github.com/roberthamel/git-repos/pkg/gitctx"
	"github.com/spf13/cobra"
)

// NewAddCmd creates a new cobra command for the add subcommand
func NewAddCmd(filename string) *cobra.Command {
	var (
		repoURL, subpath, version, localPath string
		shouldIgnore                         bool
	)
	cmd := &cobra.Command{
		Use:           "add",
		Short:         "Adds an entry to the .gitrepos file",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			gre, err := gitctx.New(repoURL, subpath, version, localPath, shouldIgnore)
			if err != nil {
				return err
			}
			gre.AddToFile(filename)
			logger.NewLog(cmd)("Added entry to .gitrepos")
			return nil
		},
	}
	cmd.Flags().StringVarP(&repoURL, "repo-url", "r", "", "The URL of the repository")
	cmd.Flags().StringVarP(&subpath, "subpath", "s", "", "The subpath of the repository")
	cmd.Flags().StringVarP(&version, "version", "v", "", "The version of the repository")
	cmd.Flags().StringVarP(&localPath, "local-path", "l", "", "The local path of the repository")
	cmd.Flags().BoolVarP(&shouldIgnore, "ignore", "i", false, "Adds the entry to the .gitignore file")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewAddCmd(filename))
}
