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

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear entries from track",
	Long:  "clear entries from track",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		tracker, err := track.NewTracker(cfg.Directory)
		if err != nil {
			logrus.Errorf("error loading tracker %s", err)
		}

		if tracker.Tracks[name] == nil {
			logrus.Errorf("error loading track %s: %s", name, err)
			return
		}

		tracker.Tracks[name].Entries = []string{}
		err = tracker.Save(name)
		if err != nil {
			logrus.Errorf("error saving track %s: %s", name, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
