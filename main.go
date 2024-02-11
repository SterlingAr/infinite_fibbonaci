package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Fibo interface {
	// Compute all the numbers up to the nth fibbonaci number and their factors and return the decimal representation
	// WARNING: memory usage may be intense for high numbers
	ComputeAll(nth int) string
	// Compute exactly the nth number and its factors
	ComputeOnly(nth int) string
}

func main() {
	arg := os.Args[1]
	maya := NewFibo()
	t1 := time.Now().UnixMicro()
	nth, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("invalid input %s %s \n", arg, err)
		os.Exit(1)
	}
	res := maya.ComputeOnly(nth)
	t2 := time.Now().UnixMicro()
	fmt.Printf("elapsed time %d microseconds (%f sec) \n", t2-t1, float64((t2-t1))/1000000)
	fmt.Printf("fibo: %s", res)

	// maya := NewFibo()
	// res := maya.ComputeOnly(144000000)
	// fmt.Printf("fibo: %s", res)
}
