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
	"path/filepath"
	kq "searchtool/k8s_utils"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

// kqCmd represents the kq command
var kqCmd = &cobra.Command{
	Use:   "kq",
	Short: "Usage: kq 'get pods --namespace kube-system'",
	Long:  `Usage: kq 'get pods --namespace kube-system'`,
	Run: func(cmd *cobra.Command, args []string) {

		if home := homedir.HomeDir(); home != "" && kubeconfig == "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
		kq.KubeQuery(kubeconfig, args)
	},
}

func init() {
	rootCmd.AddCommand(kqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kqCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
