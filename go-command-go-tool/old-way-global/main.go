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
	for i := 0; i < 10; i++ {
		result := fib.Calc(i)
		logrus.WithFields(logrus.Fields{
			"n":     i,
			"fib_n": result,
		}).Info("Fibonacci calculated")
		fmt.Printf("fib(%d) = %d\n", i, result)
	}
}
