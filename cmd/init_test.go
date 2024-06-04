package cmd

import (
	"testing"

	"github.com/roberthamel/git-repos/internal/helpers"
	"github.com/roberthamel/git-repos/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

const testFilename = "/tmp/git-repos/.gitrepos"

func TestInitCmd(t *testing.T) {
	tt := []struct {
		args []string
		out string
		preFileExists bool
		preFileExistsAndPopulated bool
	}{
		{
			args: []string{},
			out: "Initialized .gitrepos file\n",
			preFileExists: false,
			preFileExistsAndPopulated: false,
		},
		{
			args: []string{},
			out: ".gitrepos file exists and is populated. use `status` to determine if files are synchronized and up to date\n",
			preFileExists: true,
			preFileExistsAndPopulated: true,
		},
		{
			args: []string{},
			out: ".gitrepos already initialized. use `add` subcommand to add entries to the .gitrepos file\n",
			preFileExists: true,
			preFileExistsAndPopulated: false,
		},
		{
			args: []string{"--force"},
			out: "Initialized .gitrepos file\n",
			preFileExists: true,
			preFileExistsAndPopulated: true,
		},
	}
	for _, tc := range tt {
		root := NewInitCmd(testFilename)
		testhelpers.BeforeTest(t, testFilename)
		if tc.preFileExists {
			content := ""
			if tc.preFileExistsAndPopulated {
				content = "-? https://github.com/roberthamel/git-repos * main testing\n"
			}
			err := helpers.WriteFile(testFilename, content)
			assert.NoError(t, err)
		}
		testhelpers.ExecuteTest(testhelpers.ExecuteTestArgs{
			T: t,
			Cmd: root,
			Args: tc.args,
			Expected: tc.out,
		})
		assert.True(t, helpers.FileExists(testFilename))

		if tc.preFileExistsAndPopulated && len(tc.args) > 0 {
			ok, err := helpers.IsFileEmpty(testFilename)
			assert.NoError(t, err)
			assert.True(t, ok)
		}
		testhelpers.AfterTest(t, testFilename)
	}
}
