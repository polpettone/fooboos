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
	"github.com/polpettone/fooboos/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"strings"
)

func NewEditCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "edit <keyword>",
		Long:  "Open fooboos with given keyword in a text editor. If keyword is missing all fooboos will be open in a text editor.",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleEditCommand(args)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleEditCommand(args []string) (string, error) {
	fooboosFile := viper.GetString("path_to_fooboos")

	if len(args) < 1 {
		content, err := loadRaw(fooboosFile)
		if err != nil {
			return "", err
		}

		result, err := pkg.CaptureInputFromEditor(string(content))

		fooboos := &Fooboos{}
		err = yaml.Unmarshal([]byte(result), &fooboos)
		if err != nil {
			return "", err
		}

		err = writeRaw(fooboosFile, []byte(result))

		if err != nil {
			return "", err
		}

		return "saved", nil
	} else {
		keyword := args[0]

		fooboosFile := viper.GetString("path_to_fooboos")
		fooboos, err := loadFooboos(fooboosFile)

		if err != nil {
			return "", err
		}

		urls, found := fooboos.Entries[keyword]

		if !found {
			return fmt.Sprintf("no bookmark found for keyword %s", keyword), nil
		}

		content := ""
		for _, url := range urls {
			content += url + "\n"
		}

		rawResult, err := pkg.CaptureInputFromEditor(content)
		result := strings.Split(rawResult, "\n")
		fooboos.Entries[keyword] = result
		out, err := yaml.Marshal(fooboos)

		if err != nil {
			return "", err
		}

		err = writeRaw(fooboosFile, out)

		if err != nil {
			return "", err
		}

		return "done", nil
	}
}

func init() {
	editCmd := NewEditCmd()
	rootCmd.AddCommand(editCmd)
}
