// Copyright © 2018 Thomas Darimont <thomas.darimont@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the available configuration settings",
	Long:  `Lists the available configuration settings with their origin.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Using config file %s\n", viper.ConfigFileUsed())

		for k, v := range viper.AllSettings() {
			fmt.Printf("%s: %v\n", k, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
