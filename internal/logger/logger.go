package logger

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewLog(cmd *cobra.Command) func(msg any) {
	return func(msg any) {
		fmt.Fprint(cmd.ErrOrStderr(), msg, "\n")
	}
}
