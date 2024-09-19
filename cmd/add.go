/*
Copyright © 2024 Shivam Sharma <shivamsha2100@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates the task",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			log.Fatal("You must provide a task")
		}
		task := Task{Name: args[0], ID: ""}
		err := saveTasks(task)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task added: %s\n", task)
		
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
