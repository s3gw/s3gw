package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GetRootCommand returns the root cobra command to be executed
// by main.
func GetRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "s3gw",
		Short: "Start object storage gateway",
		RunE: func(cmd *cobra.Command, args []string) error {
			return root()
		},
	}
	return cmd
}

func root() error {
	fmt.Println("Hello S3 Gateway!")
	return nil
}
