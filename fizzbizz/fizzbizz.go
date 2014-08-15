package fizzbizz

// Sum of all numbers below n which are multiples of 3 or 5

func solve(n int) int {
	return solveWithStart(0, n)
}

func solveWithStart(start int, n int) int {
	sum := 0
	for i := start; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func solveGoRoutine(n int) int {
	//In theory this should give a 4x speed boost
	sum := 0
	results := make(chan int, 4)
	go func(c chan int) {
		c <- solveWithStart(0, n/4)
	}(results)
	go func(c chan int) {
		c <- solveWithStart(n/4, n/2)
	}(results)
	go func(c chan int) {
		c <- solveWithStart(n/2, 3*n/4)
	}(results)
	go func(c chan int) {
		c <- solveWithStart(3*n/4, n)
	}(results)
	sum += <-results
	sum += <-results
	sum += <-results
	sum += <-results
	return sum
}
