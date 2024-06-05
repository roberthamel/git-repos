package gitctx

import (
	"fmt"
	// "strings"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew_ValidSchema(t *testing.T) {
	gr, err := New("https://valid.url", "subpath", "version", "localPath", false)
	assert.NoError(t, err)
	assert.Equal(t, gr.String(), "-? https://valid.url subpath version localPath\n")
}

func TestValidate_EmptyRepoURL(t *testing.T) {
	gr, err := New("", "subpath", "version", "localPath", false)
	assert.Nil(t, gr)
	errs := strings.Split(err.Error(), "\n")
	assert.Equal(t, 2, len(errs))
	assert.Equal(t, "- missing required flag --repo-url", errs[0])
	assert.Equal(t, "- invalid repo URL. Must start with https:// or git@", errs[1])
}

func TestValidate_InvalidRepoURL(t *testing.T) {
	gr, err := New("invalidURL", "subpath", "version", "localPath", false)
	assert.Nil(t, gr)
	errs := strings.Split(err.Error(), "\n")
	assert.Equal(t, len(errs), 1)
	assert.Equal(t, errs[0], "- invalid repo URL. Must start with https:// or git@")
}

func TestValidate_EmptyVersion(t *testing.T) {
	gr, err := New("https://valid.url", "subpath", "", "localPath", false)
	assert.Nil(t, gr)
	errs := strings.Split(err.Error(), "\n")
	assert.Equal(t, len(errs), 1)
	assert.Equal(t, errs[0], "- missing required flag --version")
}

func TestValidate_ValidGitCtxHttps(t *testing.T) {
	gr, err := New("https://valid.url", "subpath", "version", "localPath", false)
	assert.NoError(t, err)
	assert.Equal(t, gr.String(), "-? https://valid.url subpath version localPath\n")
}

func TestValidate_ValidGitCtxGit(t *testing.T) {
	gr, err := New("git@valid.url", "subpath", "version", "localPath", true)
	assert.NoError(t, err)
	assert.Equal(t, gr.String(), "-! git@valid.url subpath version localPath\n")
}

func TestString_ShouldIgnore(t *testing.T) {
	gr, err := New("https://valid.url", "subpath", "version", "localPath", true)
	assert.NoError(t, err)
	assert.Equal(t, gr.String(), "-! https://valid.url subpath version localPath\n")
}

func TestString_ShouldNotIgnore(t *testing.T) {
	gr, err := New("https://valid.url", "subpath", "version", "localPath", false)
	assert.NoError(t, err)
	assert.Equal(t, gr.String(), "-? https://valid.url subpath version localPath\n")
}

func TestAddTimestampToString_ValidFormat(t *testing.T) {
	gr, err := New("https://valid.url", "", "version", "", false)
	assert.NoError(t, err)
	ts := time.Now().Format(time.RFC3339)
	ss, err := gr.AddTimestampToString(ts)
	assert.NoError(t, err)
	assert.Equal(t, ss, fmt.Sprintf("-? https://valid.url * version . [%s]\n", ts))
}

func TestAddTimestampToString_InvalidFormat(t *testing.T) {
	gr, err := New("https://valid.url", "subpath", "version", "localPath", false)
	assert.NoError(t, err)
	ts := time.Now()
	ss, err := gr.AddTimestampToString(ts.String())
	assert.Equal(t, ss, "")
	assert.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "timestamp does not confirm to RFC3339:"))
}

func TestAddTimestampToString_EmptyString(t *testing.T) {
	gr, err := New("https://valid.url", "subpath", "version", "localPath", false)
	assert.NoError(t, err)
	ss, err := gr.AddTimestampToString("")
	assert.Equal(t, ss, "")
	assert.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "timestamp does not confirm to RFC3339:"))
}
