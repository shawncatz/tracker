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
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/shawncatz/tracker/track"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show contents of track",
	Long:  "show contents of track",
	Args:  cobra.ExactArgs(1),
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

		t := tracker.Tracks[name]
		fmt.Printf("Name:    %s\n", t.Name)
		fmt.Printf("Created: %s\n", t.Created)
		fmt.Printf("Updated: %s\n", t.Updated)
		fmt.Println("Entries:")
		for _, e := range t.Entries {
			fmt.Printf("  - %s\n", e)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
