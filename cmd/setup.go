package cmd

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/shawncatz/tracker/config"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup usage for tracker",
	Long:  "setup usage for tracker",
	Run: func(cmd *cobra.Command, args []string) {
		cfg = config.DefaultConfig()

		if cfgFile == "" {
			cfgFile = os.Getenv("HOME") + "/.tracker.yaml"
		}

		// Create config file
		logrus.Infof("creating config file: %s", cfgFile)
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			b, err := yaml.Marshal(cfg)
			if err != nil {
				logrus.Error(err)
			}
			err = ioutil.WriteFile(cfgFile, b, 0644)
			if err != nil {
				logrus.Error(err)
			}
		}

		// Create Directory for Tracks
		logrus.Infof("creating directory: %s", cfg.Directory)
		if _, err := os.Stat(cfg.Directory); os.IsNotExist(err) {
			err := os.Mkdir(cfg.Directory, 0755)
			if err != nil {
				logrus.Error(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
