package cmd

import (
	"testing"
)

func TestRedundantBets(t *testing.T) {
	firstBet := []int{1, 2, 3}
	secondBet := []int{2, 3, 1}
	if !isRedundant(firstBet, secondBet) {
		t.Error()
	}
}

func TestNonRedundantBets(t *testing.T) {
	firstBet := []int{1, 2, 3}
	secondBet := []int{5, 3, 4}
	if isRedundant(firstBet, secondBet) {
		t.Error()
	}
}

func TestRedundantBetsDifferentSizes(t *testing.T) {
	firstBet := []int{1, 2, 3}
	secondBet := []int{1, 2, 3, 4}
	if !isRedundant(firstBet, secondBet) {
		t.Error()
	}
}

func TestNonRedundantBetsDifferentSizes(t *testing.T) {
	firstBet := []int{1, 2, 3}
	secondBet := []int{5, 3, 4, 2}
	if isRedundant(firstBet, secondBet) {
		t.Error()
	}
}
