package cmd

import (
	"reflect"
	"sort"
	"testing"
)

func TestElementsAreGeneratedInFullRange(t *testing.T) {
	expected := make([]int, 60)
	for i := range expected {
		expected[i] = i + 1
	}
	result, _ := generateSingleBet(1, 60, 60)
	sort.Ints(result)

	if !reflect.DeepEqual(result, expected) {
		t.Log("Expected: ", expected)
		t.Log("Result: ", result)
		t.Error()
	}
}

func TestElementsAreGenerated(t *testing.T) {
	minRange := 16
	maxRange := 60
	length := 15

	result, _ := generateSingleBet(minRange, maxRange, length)

	for i := range result {
		t.Log(i)
		if result[i] < minRange || result[i] > maxRange {
			t.Log("Function generated number out of range: ", result[i])
			t.Error()
		}
	}

	if len(result) != length {
		t.Log("Function generated wrong quantity of numbers: ", result)
		t.Log("Generated number of elements: ", len(result))
		t.Log("Expected number of elements: ", length)
		t.Error()
	}
}

func TestErrorRaised(t *testing.T) {
	minRange := 61
	maxRange := 60
	length := 1
	_, err := generateSingleBet(minRange, maxRange, length)
	if err == nil {
		t.Log("An error was expected to be returned")
		t.Error()
	}
}

func TestManyEqualBetsGenerated(t *testing.T) {
	qtyNumbers := 10
	config := BetConfig{
		numberOfBets: 5,
		startRange:   1,
		endRange:     10,
	}
	bets, _ := generateBets(config, qtyNumbers)

	for i := range bets {
		sort.Ints(bets[i])
	}
	for i := 1; i < len(bets); i++ {
		if !reflect.DeepEqual(bets[i], bets[i-1]) {
			t.Log("Expected: ", bets[i])
			t.Log("Result: ", bets[i-1])
			t.Error()
		}
	}
}

func TestManyDifferentBetsGenerated(t *testing.T) {
	qtyNumbers := 3
	config := BetConfig{
		numberOfBets: 5,
		startRange:   1,
		endRange:     60,
	}
	bets, _ := generateBets(config, qtyNumbers)

	t.Log("BETS: ", bets)
	for i := range bets {
		sort.Ints(bets[i])
	}
	for i := 1; i < len(bets); i++ {
		if reflect.DeepEqual(bets[i], bets[i-1]) {
			t.Log("First: ", bets[i])
			t.Log("Second: ", bets[i-1])
			t.Error()
		}
	}
}

func TestNonRedundantBetsAreRemoved(t *testing.T) {
	qtyNumbers := 2
	config := BetConfig{
		numberOfBets: 2,
		startRange:   1,
		endRange:     3,
	}
	rawBets := [][]int{[]int{1, 2}, []int{1, 2}, []int{1, 2}}
	bets, _ := removeRedundantBets(config, qtyNumbers, rawBets)

	t.Log("BETS: ", bets)
	for i := range bets {
		sort.Ints(bets[i])
	}
	for i := 1; i < len(bets); i++ {
		if reflect.DeepEqual(bets[i], bets[i-1]) {
			t.Log("First: ", bets[i])
			t.Log("Second: ", bets[i-1])
			t.Error()
		}
	}
}
