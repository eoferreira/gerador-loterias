package cmd

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func generateSingleBet(startRange int, endRange int, qtyNumbers int) ([]int, error) {
	if qtyNumbers > (endRange - startRange + 1) {
		return nil, fmt.Errorf(
			"Quantidade de números (%d) não pode ser"+
				"maior do que range desejado", qtyNumbers)
	}
	bet := make([]int, endRange-startRange+1)

	for i := range bet {
		bet[i] = i + startRange
	}
	for i := 0; i < endRange-startRange+1; i++ {
		j := rand.Intn(endRange - startRange)
		aux := bet[j]
		bet[j] = bet[i]
		bet[i] = aux
	}
	finalBet := bet[:qtyNumbers]
	sort.Ints(finalBet)
	return finalBet, nil
}

func generateBets(config BetConfig, qtyNumbersPerBet int) ([][]int, error) {
	rand.Seed(time.Now().Unix())
	bets := make([][]int, config.numberOfBets)
	for i := 0; i < config.numberOfBets; i++ {
		bet, err := generateSingleBet(config.startRange, config.endRange, qtyNumbersPerBet)
		if err != nil {
			return nil, err
		}
		bets[i] = bet
	}
	return bets, nil
}

func removeRedundantBets(config BetConfig, qtyNumbersPerBet int, bets [][]int) ([][]int, error) {
	for i := range bets {
		for j := range bets {
			if i != j {
				fmt.Sprintln(i, j)
				for isRedundant(bets[i], bets[j]) {
					newBet, err := generateSingleBet(config.startRange, config.endRange, qtyNumbersPerBet)
					fmt.Sprintln(bets[i], bets[j], newBet)
					if err != nil {
						return nil, err
					}
					bets[j] = newBet
				}
			}
		}
	}
	return bets, nil
}

func getBets(config BetConfig, qtyNumbersPerBet int) ([][]int, error) {
	rawBets, err := generateBets(config, qtyNumbersPerBet)
	if err != nil {
		return nil, err
	}
	finalBets, err := removeRedundantBets(config, qtyNumbersPerBet, rawBets)
	if err != nil {
		return nil, err
	}
	return finalBets, nil
}
