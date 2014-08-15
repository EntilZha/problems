package fizzbizz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	assert.Equal(t, 23, solve(10))
}
func TestSolveGoRoutine(t *testing.T) {
	assert.Equal(t, 23, solveGoRoutine(10))
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(10000000)
	}
}
func BenchmarkGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveGoRoutine(10000000)
	}
}
