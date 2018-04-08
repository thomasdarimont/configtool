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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Reads the value of the given setting from the detected sources",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "Setting name missing")
			os.Exit(1)
		}

		setting := args[0]

		if value := viper.GetString(setting); value != "" {
			fmt.Println(value)
			return
		}

		if def, _ := cmd.Flags().GetString("default"); def != "" {
			fmt.Println(def)
			return
		}

		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("default", "d", "", "Default value for the configuration setting")
}
