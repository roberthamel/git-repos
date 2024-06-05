package cmd

import (
	"errors"
	"testing"

	"github.com/roberthamel/git-repos/internal/helpers"
	"github.com/roberthamel/git-repos/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestAddCmd(t *testing.T) {
	tt := []struct {
		args     []string
		err      error
		out      string
		expected string
	}{
		{
			args:     []string{"-r", "https://github.com", "-v", "main"},
			err:      nil,
			out:      "Added entry to .gitrepos\n",
			expected: "-? https://github.com * main .",
		},
		{
			args:     []string{"-r", "https://github.com", "-v", "main", "-i"},
			err:      nil,
			out:      "Added entry to .gitrepos\n",
			expected: "-! https://github.com * main .",
		},
		{
			args:     []string{"-r", "https://github.com", "-s", "/exampleA", "-v", "v1.0.0", "-l", "/exampleB"},
			err:      nil,
			out:      "Added entry to .gitrepos\n",
			expected: "-? https://github.com exampleA v1.0.0 exampleB",
		},
		{
			args:     []string{"-r", "https://github.com", "-s", "exampleA", "-v", "v1.0.0", "-i", "-l", "exampleB"},
			err:      nil,
			out:      "Added entry to .gitrepos\n",
			expected: "-! https://github.com exampleA v1.0.0 exampleB",
		},
		{
			args: []string{"-v", "main"},
			err:  errors.New("- missing required flag --repo-url\n- invalid repo URL. Must start with https:// or git@"),
		},
		{
			args: []string{"-r", "https://github.com"},
			err: errors.New("- missing required flag --version"),
		},
		{
			args: []string{"-r", "git@github.com", "-v", "v1.0.0"},
			err: nil,
			out: "Added entry to .gitrepos\n",
			expected: "-? git@github.com * v1.0.0 .",
		},
		{
			args: []string{"-r", "intentionally-incorrect", "-v", "v1.0.0"},
			err: errors.New("- invalid repo URL. Must start with https:// or git@"),
		},
	}
	for _, tc := range tt {
		root := NewAddCmd(testFilename)
		testhelpers.BeforeTest(t, testFilename)
		assert.NoError(t, helpers.WriteFile(testFilename, ""))
		testhelpers.ExecuteTest(testhelpers.ExecuteTestArgs{
			T:           t,
			Cmd:         root,
			Args:        tc.args,
			Expected:    tc.out,
			ExpectedErr: tc.err,
		})
		result, err := helpers.ReadLastLineOfFile(testFilename)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, result)
		testhelpers.AfterTest(t, testFilename)
	}
}
