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
	"log"
	"strconv"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		status,_ := cmd.Flags().GetBool("float")
		if status {
			floatAdd(args)
		} else {
			intAdd(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().BoolP("float","f",false,"add floating numbers")
}

func intAdd(args []string){
	var sum int
	for _, num  := range args{
		temp, err := strconv.Atoi(num)
		if err != nil{
			//log.Fatal()
			log.Fatal("what's fuck, please input a int type.")
			//panic(err)

		}
		sum = sum + temp
	}
	fmt.Printf("args: %v, sum:%v\n",args, sum)
}

func floatAdd(args []string){
	var sum float64
	for _, num  := range args{
		temp, err := strconv.ParseFloat(num,64)
		if err != nil{
			//log.Fatal()
			log.Fatal("what's fuck, please input a int type.")
			//panic(err)
		}
		sum = sum + temp
	}
	fmt.Printf("args: %v, sum:%v\n",args, sum)
}
