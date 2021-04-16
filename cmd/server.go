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
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/feiwang137/athena/pkg/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(" listenaddress: %v\n ruleconfigPath: %v\n promtoolPath: %v\n", listenAddress, ruleConfigPath, promToolPath)
		server.PromServer()
	},
}


var (
	//listenAddress  string
	//promToolPath   string
	serverConfigPath string
)

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//serverCmd.Flags().StringVar(&listenAddress, "listen_address", "0.0.0.0:8080", "listen address and ports.")
	//serverCmd.Flags().StringVar(&promToolPath, "promtool_path", "/usr/local/bin/", "promtool path.")

	// 拿到配置文件后，需要一个配置文件解析模块，init() 指针方式
	serverCmd.Flags().StringVar(&serverConfigPath, "serverConfigPath", "/Users/feiwang/prom-data/athena.yml", "server Config Path")
}
