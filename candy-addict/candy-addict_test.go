package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCase(t *testing.T) {
	allowances := []int{5, 5, 5, 5, 5}
	candyCosts := []int{3, 3, 3, 3, 3}
	days := []int{1, 2, 3, 4, 5}
	answers := []int{2, 1, 6, 2, 7}
	for i := range answers {
		assert.Equal(t, answers[i], solveCase(allowances[i], candyCosts[i], days[i]))
	}
}

func TestComplexCase(t *testing.T) {
	allowances := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	candyCosts := []int{1, 2, 3, 999999998, 999999999}
	days := []int{342568368, 560496730, 586947396, 386937583, 609483745}
	answers := []int{342568367000000000, 60496729000000000, 253614062000000001, 773875166, 609483745}
	for i := range answers {
		assert.Equal(t, answers[i], solveCase(allowances[i], candyCosts[i], days[i]), "Testing case "+string(i))
	}
}

func BenchmarkGoRoutine(b *testing.B) {
	for iBench := 0; iBench < b.N; iBench++ {
		var wg sync.WaitGroup
		allowances := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
		candyCosts := []int{1, 2, 3, 999999998, 999999999}
		days := []int{342568368, 560496730, 586947396, 386937583, 609483745}
		answers := []int{342568367000000000, 60496729000000000, 253614062000000001, 773875166, 609483745}
		for i := range answers {
			wg.Add(1)
			go solveCaseGoRoutine(&wg, allowances[i], candyCosts[i], days[i])
		}
		wg.Wait()
	}
}

func BenchmarkSequential(b *testing.B) {
	for iBench := 0; iBench < b.N; iBench++ {
		allowances := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
		candyCosts := []int{1, 2, 3, 999999998, 999999999}
		days := []int{342568368, 560496730, 586947396, 386937583, 609483745}
		answers := []int{342568367000000000, 60496729000000000, 253614062000000001, 773875166, 609483745}
		for i := range answers {
			solveCase(allowances[i], candyCosts[i], days[i])
		}
	}
}
