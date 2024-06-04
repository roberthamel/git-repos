package cmd

import (
	"fmt"
	"testing"

	"github.com/roberthamel/git-repos/internal/helpers"
	"github.com/roberthamel/git-repos/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	testhelpers.BeforeTest(t, testFilename)
	assert.False(t, helpers.FileExists(testFilename))
	root := NewRootCmd()
	testhelpers.ExecuteTest(testhelpers.ExecuteTestArgs{
		T:        t,
		Cmd:      root,
		Args:     []string{},
		Expected: fmt.Sprintf("%s\n\n", rootShortHelp),
	})
	testhelpers.AfterTest(t, testFilename)
}
