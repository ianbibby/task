/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			n, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Invalid task id:", arg)
				os.Exit(1)
			}
			ids = append(ids, n)
		}

		tasks, err := App.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Invalid task id: %d\n", id)
				continue
			}

			task := tasks[id-1]
			err := App.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark %q as completed.  Error: %v", task.Val, err)
				continue
			}

			fmt.Printf("Marked %q as completed.\n", task.Val)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
