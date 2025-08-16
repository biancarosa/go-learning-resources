package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Fibonacci struct{}

func NewFibonacci() *Fibonacci {
	return &Fibonacci{}
}

func (f *Fibonacci) Calc(n int) int {
	if n <= 1 {
		return n
	}
	logrus.WithField("n", n).Debug("Calculating fib(n-1) + fib(n-2)")
	return f.Calc(n-1) + f.Calc(n-2)
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	fib := NewFibonacci()
	nums := make([]int, 10)
	for i := range nums {
		result := fib.Calc(i)
		logrus.WithFields(logrus.Fields{
			"n":     i,
			"fib_n": result,
		}).Info("Fibonacci calculated")
		fmt.Printf("fib(%d) = %d\n", i, result)
	}

	// go vet
	x := 5
	fmt.Printf("The value is %s\n", x)

}

// staticcheck
func unusedFunc() {
	fmt.Println("I am never called")
}
