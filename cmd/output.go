package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func printOutput(bets [][]int) {
	fmt.Printf("DEZENAS GERADAS:\n")
	for _, value := range bets {
		fmt.Printf("%v\n", value)
	}
}

func writeCsv(betData [][]int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Erro ao salvar arquivo com apostas", err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range betData {
		output := strings.Fields(strings.Trim(fmt.Sprint(value), "[]"))
		err := writer.Write(output)
		if err != nil {
			log.Fatal("Erro ao salvar arquivo com apostas", err)
		}
	}
}
