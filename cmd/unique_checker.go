package cmd

func contains(bet []int, number int) bool {
	for _, a := range bet {
		if a == number {
			return true
		}
	}
	return false
}

func isRedundant(firstBet []int, secondBet []int) bool {
	baseBet := firstBet
	comparisonBet := secondBet
	if len(firstBet) < len(secondBet) {
		baseBet = secondBet
		comparisonBet = firstBet
	}
	for _, number := range comparisonBet {
		if !contains(baseBet, number) {
			return false
		}
	}
	return true
}
