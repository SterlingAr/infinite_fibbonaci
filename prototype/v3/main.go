package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func main() {
	arg := os.Args[1]
	nth, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("invalid input %s %s \n", arg, err)
		os.Exit(1)
	}
	t1 := time.Now().UnixMicro()
	res := compute(nth)
	t2 := time.Now().UnixMicro()
	fmt.Printf("elapsed time %d microseconds (%f sec) \n", t2-t1, float64((t2-t1))/1000000)
	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Printf("failed to create file %s \n",  err)
		os.Exit(1)
	}
	fmt.Fprint(f, mayanLongCountDec(res).String())
}

func compute(n int) []byte {
	previous := []byte{1}
	current := []byte{1}

	for i := 1; i < n-1; i++ {
		next := addColumnsInPlace(previous, current)

		// Update for the next iteration
		previous, current = current, next

		// Optionally, print the current step
		fmt.Printf("Step %d: %v\n", i+1, mayanLongCountDec(current).String())
	}
	return current
}

func dec(n []byte) *big.Int {
	power := len(n) - 1
	res := new(big.Int)
	for i := 0; i <= power; i++ {
		z := new(big.Int).Exp(big.NewInt(20), big.NewInt(int64(i)), nil)
		z = z.Mul(z, big.NewInt(int64(n[i])))
		res.Add(res, z)
	}
	return res
}

// mayanLongCountDec converts a byte slice representing a Mayan Long Count date to a *big.Int
func mayanLongCountDec(n []byte) *big.Int {
	res := new(big.Int)
	placeValue := big.NewInt(1)
	base := int64(20)

	for i, val := range n {
		z := new(big.Int).Mul(placeValue, big.NewInt(int64(val)))
		res.Add(res, z)

		// Update placeValue for next iteration
		if i == 0 {
			base = 18 // Second position uses base-18
		} else {
			base = 20 // Reset to base-20 for subsequent positions
		}
		placeValue.Mul(placeValue, big.NewInt(base))
	}

	return res
}

// addColumnsInPlace adds two slices of byte with carry-over, updating the first slice in place.
func addColumnsInPlace(a, b []byte) []byte {
	maxLen := len(a)
	if len(b) > maxLen {
		// Extend 'a' to the length of 'b' if necessary
		extended := make([]byte, len(b))
		copy(extended, a)
		a = extended
		maxLen = len(b)
	}

	carry := byte(0)
	for i := 0; i < maxLen; i++ {
		sum := carry
		if i < len(a) {
			sum += a[i]
		}
		if i < len(b) {
			sum += b[i]
		}

		if sum >= 20 {
			carry = 1
			sum -= 20
		} else {
			carry = 0
		}

		a[i] = sum
	}

	if carry > 0 {
		a = append(a, carry)
	}

	return a
}
