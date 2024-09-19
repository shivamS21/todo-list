/*
Copyright © 2024 Shivam Sharma <shivamsha2100@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("Delete task: ")
		log.SetFlags(0)
		if len(args[0]) < 1 {
			log.Fatal("Provide valid task to be deleted")
		}
		err := deleteTaskByID(args[0])

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Task Deleted successfully!")

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
