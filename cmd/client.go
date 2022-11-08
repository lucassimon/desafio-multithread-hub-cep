/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"time"

	"github.com/lucassimon/desafio-multithread-hub-cep/internal/dto"
	"github.com/lucassimon/desafio-multithread-hub-cep/internal/providers"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
		via_cep_chan := make(chan *dto.CepOutput)
		api_cep_chan := make(chan *dto.CepOutput)
		postmon_chan := make(chan *dto.CepOutput)

		viacep := providers.NewViaCep()
		apicep := providers.NewApiCep()
		postmon := providers.NewPostmonCep()
		go viacep.Search(via_cep_chan, "05477-902")
		go apicep.Search(api_cep_chan, "05477-902")
		go postmon.Search(postmon_chan, "05477902")

		select {
		case result := <-via_cep_chan:
			fmt.Println("ViaCep", result)
		case result := <-api_cep_chan:
			fmt.Println("ApiCep", result)
		case result := <-postmon_chan:
			fmt.Println("Postmon", result)
		case <-time.After(time.Second * 1):
			fmt.Println("Não conseguimos buscar o resultado")
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
