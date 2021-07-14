package main

import "fmt"

func main() {
	minAmount := 5000
	purchase := 15_000
	percentageCashback :=15
	limitCashback := 2000
	cashback := 0
	const fullPercent int = 100

	if purchase >= minAmount{
		cashback = purchase * percentageCashback / fullPercent
		fmt.Println("ffffff")
	}

	if cashback > limitCashback{
		cashback = limitCashback
	}
	fmt.Println("gggg")
}
