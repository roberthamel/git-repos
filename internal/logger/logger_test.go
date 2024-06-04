package logger

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {
	cmd := &cobra.Command{}
	buffer := new(bytes.Buffer)
	cmd.SetErr(buffer)
	log := NewLog(cmd)
	log("test message")
	expected := "test message\n"
	got := buffer.String()
	assert.Equal(t, expected, got)
}
