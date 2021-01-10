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
	"sort"
)


func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all fooboo keywords",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleListCommand()

			if err != nil {
				fmt.Println(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleListCommand() (string, error) {
	fooboosFile := viper.GetString(FooboosFile)
	fooboos, err := loadFooboos(fooboosFile)

	if err != nil {
		fmt.Printf("%v", err)
		return "", err
	}


	output := ""

	var keywords []string

	for k, _ := range fooboos.Entries {
		keywords = append(keywords, k)
	}

	sort.Strings(keywords)

	for _, k := range keywords {
		output += k + "\n"
	}

	return output, nil
}

func init() {
	listCmd := NewListCmd()
	rootCmd.AddCommand(listCmd)
}
