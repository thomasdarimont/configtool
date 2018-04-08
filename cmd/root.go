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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var appName string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "configtool",
	Short: "configtool reads configuration from multiple sources.",
	Long:  `configtool reads configuration from multiple sources like parameters, files and env variables.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&appName, "app", "a", "", "app name used to derive configuration file search paths (optional)")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config", "config name used to derive the configuration file name, defaults to 'config'")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if appName == "" {
		appName = os.Getenv("CONFIGTOOL_APP")
	}

	if cfgFile == "" {
		cfgFile = os.Getenv("CONFIGTOOL_CONFIG")
	}

	if strings.Contains(cfgFile, "/") {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(cfgFile)                 // name of config file (without extension)
		viper.AddConfigPath("./." + appName)         // optionally look for config in the working directory
		viper.AddConfigPath("$HOME/." + appName)     // call multiple times to add many search paths
		viper.AddConfigPath("/etc/" + appName + "/") // path to look for the config file in
	}

	viper.SetEnvPrefix(appName)
	viper.BindEnv("app")
	viper.BindEnv("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not find config file, err: %v\n", err)
	}
}
