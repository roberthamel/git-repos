package gitctx

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// GitCtx represents a single entry in the .gitrepos file
type GitCtx struct {
	RepoURL      string
	Subpath      string
	Version      string
	LocalPath    string
	ShouldIgnore bool
}

// New creates a new GitRepo object
func New(repoURL, subpath, version, localPath string, ignore bool) (*GitCtx, error) {
	gctx := &GitCtx{
		RepoURL:      repoURL,
		Subpath:      subpath,
		Version:      version,
		LocalPath:    localPath,
		ShouldIgnore: ignore,
	}
	errs := gctx.validate()
	if len(errs) > 0 {
		errStrings := make([]string, len(errs))
		for i, err := range errs {
			errStrings[i] = fmt.Sprintf("- %s", err.Error())
		}
		return nil, errors.New(strings.Join(errStrings, "\n"))
	}
	return gctx, nil
}

// String returns a string representation of the GitRepo object
func (gr *GitCtx) String() string {
	var ignore string
	if gr.ShouldIgnore {
		ignore = "!"
	} else {
		ignore = "?"
	}
	return fmt.Sprintf("-%s %s %s %s %s\n", ignore, gr.RepoURL, gr.Subpath, gr.Version, gr.LocalPath)
}

// AddTimestampToString appends a timestamp to the end of the string and validates
// that the supplied timestamp is in the correct format (RFC3339)
func (gr *GitCtx) AddTimestampToString(timestamp string) (string, error) {
	_, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return "", fmt.Errorf("timestamp does not confirm to RFC3339: %v", err)
	}
	ss := strings.TrimRight(gr.String(), "\n")
	return fmt.Sprintf("%s [%s]\n", ss, time.Now().Format(time.RFC3339)), nil
}

// AddToFile appends the GitRepo object to the end of the file
func (gr *GitCtx) AddToFile(filename string) error {
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
func (gr *GitCtx) validate() []error {
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
