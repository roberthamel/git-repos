package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/roberthamel/git-repos/internal/logger"
	"github.com/spf13/cobra"
)

// NewAddCmd creates a new cobra command for the add subcommand
func NewAddCmd(filename string) *cobra.Command {
	var (
		gitRepoEntry                         GitRepo
		repoURL, subpath, version, localPath string
		shouldIgnore                         bool
	)
	cmd := &cobra.Command{
		Use:           "add",
		Short:         "Adds an entry to the .gitrepos file",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			gitRepoEntry = *NewGitRepo(repoURL, subpath, version, localPath, shouldIgnore)
			errs := gitRepoEntry.Validate()
			if len(errs) > 0 {
				errStrings := make([]string, len(errs))
				for i, err := range errs {
					errStrings[i] = err.Error()
				}
				return errors.New(strings.Join(errStrings, "\n"))
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			gitRepoEntry.AddToFile(filename)
			logger.NewLog(cmd)("Added entry to .gitrepos")
		},
	}
	cmd.Flags().StringVarP(&repoURL, "repo-url", "r", "", "The URL of the repository")
	cmd.Flags().StringVarP(&subpath, "subpath", "s", "", "The subpath of the repository")
	cmd.Flags().StringVarP(&version, "version", "v", "", "The version of the repository")
	cmd.Flags().StringVarP(&localPath, "local-path", "l", "", "The local path of the repository")
	cmd.Flags().BoolVarP(&shouldIgnore, "ignore", "i", false, "Adds the entry to the .gitignore file")
	return cmd
}

// GitRepo represents a single entry in the .gitrepos file
type GitRepo struct {
	RepoURL      string
	Subpath      string
	Version      string
	LocalPath    string
	ShouldIgnore bool
}

// NewGitRepo creates a new GitRepo object
func NewGitRepo(repoURL, subpath, version, localPath string, ignore bool) *GitRepo {
	return &GitRepo{
		RepoURL:      repoURL,
		Subpath:      subpath,
		Version:      version,
		LocalPath:    localPath,
		ShouldIgnore: ignore,
	}
}

// String returns a string representation of the GitRepo object
func (gr GitRepo) String() string {
	var ignore string
	if gr.ShouldIgnore {
		ignore = "!"
	} else {
		ignore = "?"
	}
	// return fmt.Sprintf("-%s %s %s %s %s [%s]\n", ignore, gr.RepoURL, gr.Subpath, gr.Version, gr.LocalPath, time.Now().Format(time.RFC3339))
	return fmt.Sprintf("-%s %s %s %s %s\n", ignore, gr.RepoURL, gr.Subpath, gr.Version, gr.LocalPath)
}

// AddToFile appends the GitRepo object to the end of the file
func (gr *GitRepo) AddToFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(gr.String())
	if err != nil {
		return err
	}
	return nil
}

// Validate checks if the GitRepo object has all required fields
func (gr *GitRepo) Validate() []error {
	err := make([]error, 0)
	if gr.RepoURL == "" {
		err = append(err, errors.New("missing required flag --repo-url"))
	}
	if !strings.HasPrefix(gr.RepoURL, "https://") && !strings.HasPrefix(gr.RepoURL, "git@") {
		err = append(err, errors.New("invalid repo URL. Must start with https:// or git@"))
	}
	if gr.Subpath == "" {
		gr.Subpath = "*"
	}
	gr.Subpath = strings.TrimPrefix(gr.Subpath, "/")
	if gr.Version == "" {
		err = append(err, errors.New("missing required flag --version"))
	}
	if gr.LocalPath == "" {
		gr.LocalPath = "."
	}
	gr.LocalPath = strings.TrimPrefix(gr.LocalPath, "/")
	if len(err) > 0 {
		return err
	}
	return []error{}
}

func init() {
	rootCmd.AddCommand(NewAddCmd(filename))
}
