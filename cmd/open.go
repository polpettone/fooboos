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
	"os/exec"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "keyword",
	Run: func(cmd *cobra.Command, args []string) {
		handleOpenCommand(args)
	},
}

func handleOpenCommand(args []string) {

	if len(args) < 1 {
		fmt.Printf("missing fooboo keyword")
		return
	}

	keyword := args[0]

	fooboosFile := viper.GetString("path_to_fooboos")
	fooboos, err := loadFooboos(fooboosFile)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	urls, found := fooboos.Entries[keyword]

	if !found {
		fmt.Printf("no bookmark found for keyword %s", keyword)
		return
	}

	for _, url := range urls {
		cmd := exec.Command("firefox", url)
		cmd.Start()
	}

}

func init() {
	rootCmd.AddCommand(openCmd)
}
