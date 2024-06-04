package cmd

import (
	"github.com/spf13/cobra"
)

var rootShortHelp = "Track external version controlled files as an alternative to git submodules"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git-repos",
		Short: rootShortHelp,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
			}
			return nil
		},
	}
	return cmd
}

var rootCmd = NewRootCmd()

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
