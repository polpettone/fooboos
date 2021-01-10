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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func SearchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "search for a fooboo by keyword",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleSearchCommand(args)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleSearchCommand(args []string) (string, error) {

	if len(args) != 1 {
		return fmt.Sprintf("you need exactly one argument -> search query"), nil
	}

	fooboosFile := viper.GetString("path_to_fooboos")
	fooboos, err := loadFooboos(fooboosFile)

	if err != nil {
		return "", err
	}

	result := fooboos.search(args[0])

	output := ""
	if result == nil {
		output = "nix gefunden"
	} else {
		for _, v := range result {
			output += v + "\n"
		}
	}

	return output, nil
}

func init() {
	searchCmd := SearchCmd()
	rootCmd.AddCommand(searchCmd)
}
