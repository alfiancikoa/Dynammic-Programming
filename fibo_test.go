package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fibonacci(t *testing.T) {
	input := 10
	output := Fibonacci(input)
	assert.Equal(t, 55, output, "Expected 55")
}

// go test -v -cover

func BenchmarkFibo(b *testing.B) {
	Fibonacci(100)
}

// go test -bench=.
