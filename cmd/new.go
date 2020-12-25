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
	"gopkg.in/yaml.v2"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new <new keyword> <url>",
		Short: "add a new fooboo",
		Run: func(cmd *cobra.Command, args []string) {
			//TODO: check args lenght
			stdout, err := handleNewCommand(args)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleNewCommand(args []string) (string, error) {

	if len(args) != 2 {
		return fmt.Sprintf("too few arguments. You need a new keyword and a url"), nil
	}

	newKeyWord := args[0]
	url := args[1]

	fooboosFile := viper.GetString("path_to_fooboos")
	fooboos, err := loadFooboos(fooboosFile)

	if err != nil {
		return "", err
	}

	for k, e := range fooboos.Entries {
		if k == newKeyWord {
			//TODO: better output format
			o := fmt.Sprintf("keyword already exists with following entries: %v", e )
			return o, nil
		}
	}

	fooboos.Entries[newKeyWord] = []string{url}
	out, err := yaml.Marshal(fooboos)
	if err != nil {
		return "", err
	}

	err = writeRaw(fooboosFile, out)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("added new fooboo with keyword: %s", newKeyWord), nil
}

func init() {
	newCmd := NewCmd()
	rootCmd.AddCommand(newCmd)
}
