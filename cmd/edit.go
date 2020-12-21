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
)


func NewEditCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "edit",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleEditCommand()

			if err != nil {
				fmt.Println(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleEditCommand() (string, error) {
	fooboosFile := viper.GetString("path_to_fooboos")
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
}

func init() {
	editCmd := NewEditCmd()
	rootCmd.AddCommand(editCmd)
}
