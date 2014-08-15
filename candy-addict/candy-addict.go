package main

import "sync"

// This code solves the problem statement found here
// http://community.topcoder.com/stat?c=problem_statement&pm=13345

func main() {
}

func solveCase(allowance int, candyCost int, nDays int) int {
	money := 0
	candies := 0
	for day := 1; day <= nDays; day++ {
		money += allowance
		if candies == 0 {
			nCandies := money / candyCost
			money -= nCandies * candyCost
			candies += nCandies
		}
		candies -= 1
	}
	return money
}

func solveCaseGoRoutine(wg *sync.WaitGroup, allowance int, candyCost int, nDays int) int {
	defer wg.Done()
	return solveCase(allowance, candyCost, nDays)
}
