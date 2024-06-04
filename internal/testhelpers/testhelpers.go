package testhelpers

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/roberthamel/git-repos/internal/helpers"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type ExecuteTestArgs struct {
	T           *testing.T
	Cmd         *cobra.Command
	Args        []string
	Expected    string
	ExpectedErr error
}

func ExecuteTest(a ExecuteTestArgs) {
	a.T.Helper()
	buf := new(bytes.Buffer)
	a.Cmd.SetOut(buf)
	a.Cmd.SetErr(buf)
	a.Cmd.SetArgs(a.Args)
	err := a.Cmd.Execute()
	assert.Equal(a.T, a.ExpectedErr, err)
	assert.Equal(a.T, a.Expected, buf.String())
}

func AfterTest(t *testing.T, filename string) {
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
	assert.True(t, !helpers.FileExists(filename))
}

func BeforeTest(t *testing.T, filename string) {
	if helpers.FileExists(filename) {
		os.Remove(filename)
	}
	dirname := filepath.Dir(filename)
	err := os.Remove(dirname)
	assert.NoError(t, err)
	err = os.MkdirAll(dirname, 0755)
	assert.NoError(t, err)
	assert.True(t, helpers.FileExists(dirname))
	assert.False(t, helpers.FileExists(filename))
}
