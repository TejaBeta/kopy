/*
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
	kopy "kopy/internal"
	"kopy/internal/options"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	context     string
	kopyOptions = options.GetKopyOptions(context)
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kopy",
	Short: "kopy is a kubectl plugin to copy resources",
	Long: `kopy is a kubectl plugin to copy resources from one context to another context

kopy is a kubectl plugin or a cli to copy K8s resources
from a particular namespace from the current context to
another context. However, the context switching is kind
of opinionated, as the tool requires all the contexts 
within the same config.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		kopy.Kopy(kopyOptions)
	},
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

	rootCmd.Flags().StringVar(&kopyOptions.Namespace, "ns", "", "Namespace within the current context")
	rootCmd.Flags().StringVarP(&context, "context", "c", "", "Context name to copy resources into(required)")
	rootCmd.Flags().BoolVar(&kopyOptions.IsAll, "a", false, "All the resources within the namespace")
	rootCmd.MarkFlagRequired("ns")
	rootCmd.MarkFlagRequired("context")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kopy" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kopy")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
