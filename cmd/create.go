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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/shawncatz/tracker/track"
)

// createCmd represents the add command
var createCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "create new track",
	Long:  "create new track",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		tracker, err := track.NewTracker(cfg.Directory)
		if err != nil {
			logrus.Errorf("error loading tracker %s", err)
		}

		if tracker.Tracks[name] == nil {
			err = tracker.NewTrack(name)
			if err != nil {
				logrus.Errorf("error creating track %s: %s", name, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
