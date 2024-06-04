/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/roberthamel/git-repos/internal/helpers"
	"github.com/roberthamel/git-repos/internal/logger"
	"github.com/spf13/cobra"
)

const filename = "./.gitrepos"


func NewInitCmd(filename string) *cobra.Command {
	force := false
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Adds a new `.gitrepos` file to the current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.NewLog(cmd)
			if helpers.FileExists(filename) && !force {
				fileIsEmpty, err := helpers.IsFileEmpty(filename)
				if err != nil {
					return fmt.Errorf("error checking if .gitrepos file is empty: %w", err)
				}
				if !fileIsEmpty {
					log(".gitrepos file exists and is populated. use `status` to determine if files are synchronized and up to date")
					return nil
				}
				log(".gitrepos already initialized. use `add` subcommand to add entries to the .gitrepos file")
				return nil
			}
			err := helpers.WriteFile(filename, "")
			if err != nil {
				return err
			}
			log("Initialized .gitrepos file")
			return nil
		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Write a new .gitrepos file")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewInitCmd(filename))
}
