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
	"log"
	"os"
	"github.com/feiwang137/athena/pkg/agent"
	"github.com/spf13/cobra"
	"github.com/feiwang137/athena/pkg/utils"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("init called, promehtues url:%v and init db\n",prometheusServer)
		err := agent.InitAthena(prometheusServer)
		if err != nil{
			log.Println(err)
			os.Exit(1)
		}
		err = utils.GenServerConfig(prometheusServer,sqliteDbPath,ruleConfigPath,listenAddress,promToolPath,athenaConfigPath)
		if err !=nil{
			log.Fatalln("generate athena.yaml faild.")
		}

	},
}

var (
	prometheusServer string
	sqliteDbPath string
	ruleConfigPath string
	listenAddress  string
	promToolPath   string
	athenaConfigPath   string
)





func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCmd.Flags().StringVar(&prometheusServer,"prometheus_url","http://127.0.0.1:9090"," prometheus url")
	initCmd.Flags().StringVar(&sqliteDbPath,"sqliteDBPath","/Users/feiwang/prom-data/athena.db","sqliteDB Path")
	serverCmd.Flags().StringVar(&ruleConfigPath, "rule_config_path", "/Users/feiwang/prom-data/rules.yml", "rule config path.")
	serverCmd.Flags().StringVar(&listenAddress, "listen_address", "0.0.0.0:8080", "listen address and ports.")
	serverCmd.Flags().StringVar(&promToolPath, "promtool_path", "/usr/local/bin/promtool", "promtool path.")
	serverCmd.Flags().StringVar(&athenaConfigPath, "athena_config_path", "/Users/feiwang/prom-data/athena.yml", "athena config path.")

}
