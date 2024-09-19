/*
Copyright © 2024 Shivam Sharma <shivamsha2100@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("    Task List\n-----------------")
		tasks, err := loadTasks()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("ID\tTask\n-- \t----\n")
		for _, task:= range tasks {
			fmt.Printf("%s\t%s\n\n", task.ID, task.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
