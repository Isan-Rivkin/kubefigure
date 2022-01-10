/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)



// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "fetch input from data sources",
	Long: `use subcommand
$input terraform --backed-type=s3 --backend-s3-bucket=bucket-name --backend-key=some/key/state --backend-region=us-east-1 --state-value-path=outputs.alb_sg --status-only
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("input called")
	},
}

func init() {
	rootCmd.AddCommand(inputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
