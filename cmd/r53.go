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
	awsu "searchtool/aws_utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// r53Cmd represents the r53 command
var r53Cmd = &cobra.Command{
	Use:   "r53",
	Short: "Query route53 records values",
	Long: `Query route53 records values

	For example:

	st r53 -r '*.my.sub.domain.com'
`,
	Run: func(cmd *cobra.Command, args []string) {
		if debug == true {
			log.SetLevel(log.DebugLevel)
		}

		if recordInput == "" {
			log.Error("query must not be empty use -r `my.domain.com`")
			return
		}
		api := awsu.NewRoute53Api()
		result, err := api.GetRecordSetAliases(recordInput)
		if err != nil {
			log.WithError(err).Error("failed")
			return
		}
		if result == nil {
			log.Error("no result found")
			return
		}
		result.PrintTable()
	},
}

func init() {
	rootCmd.AddCommand(r53Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// r53Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// r53Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
