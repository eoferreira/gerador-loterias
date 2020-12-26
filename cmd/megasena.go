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
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// megasenaCmd represents the megasena command
var megasenaCmd = &cobra.Command{
	Use:   "megasena",
	Short: "Geração de dezenas da mega sena",
	Run: func(cmd *cobra.Command, args []string) {
		getMegasenaBets(numberOfBets, 1, 60, qtyNumbersPerBet)
	},
}

func init() {
	rootCmd.AddCommand(megasenaCmd)
}

func getMegasenaBets(numberOfBets int, startRange int, endRange int, qtyNumbersPerBet int) {
	fmt.Println("Gerando dezenas da mega sena")
	config := BetConfig{
		numberOfBets: numberOfBets,
		startRange:   startRange,
		endRange:     endRange,
	}
	generatedBets, err := getBets(config, qtyNumbersPerBet)
	if err != nil {
		log.Fatal("Algo de errado aconteceu", err)
	}
	printOutput(generatedBets)
	writeCsv(generatedBets, "mega.csv")
}
