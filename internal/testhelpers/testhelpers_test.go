package testhelpers

import (
	"os"
	"testing"

	"github.com/roberthamel/git-repos/internal/helpers"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestExecuteTest(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "A test command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("test output")
		},
	}
	args := ExecuteTestArgs{
		T: t,
		Cmd: cmd,
		Args: []string{},
		Expected: "test output\n",
	}
	ExecuteTest(args)
}

func TestAfterTest(t *testing.T) {
	filename := "testfile"
	_, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	AfterTest(t, filename)
	assert.False(t, helpers.FileExists(filename))
}

func TestBeforeTest(t *testing.T) {
	filename := "/tmp/testdir/testfile"
	BeforeTest(t, filename)
	assert.True(t, helpers.FileExists("/tmp/testdir"))
	assert.False(t, helpers.FileExists(filename))
}
