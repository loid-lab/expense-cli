package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an Expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Expense deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
