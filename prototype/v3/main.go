package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	// Number of iterations
	// Initial value Q
	// Initial value W
	// Update for the next iteration
	// Print current iteration result
	t1 := time.Now().UnixMicro()
	res := compute(144000)
	t2 := time.Now().UnixMicro()
	fmt.Printf("elapsed time %d microseconds (%f sec) \n", t2-t1, float64((t2-t1))/1000000)
	fmt.Printf("fibo: %s\n", res)
}

func compute(n int) string {

	previous := []byte{1}
	current := []byte{1}

	for i := 1; i < n-1; i++ {
		next := addColumns(previous, current)

		previous = current
		current = next
		// fmt.Printf("Step %d: %v\n", i+1, current)
	}
	return dec(current).String() 
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

// addColumns adds two slices of byte with carry-over.
func addColumns(a, b []byte) []byte {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	result := make([]byte, maxLen)
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

		result[i] = sum
	}

	if carry > 0 {
		result = append(result, carry)
	}

	return result
}
