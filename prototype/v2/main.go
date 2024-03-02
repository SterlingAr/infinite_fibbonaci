package main

import (
	"fmt"
	"strings"
)

func main() {
	Q := []int{1}
	W := []int{1}
	var progression []string

	progression = append(progression, formatNumber(Q), formatNumber(W))

	current := W
	previous := Q

	for i := 0; i < 20; i++ { // You can adjust the number of iterations to generate more elements in the progression
		next := addNumbers(previous, current)
		progression = append(progression, formatNumber(next))

		previous = current
		current = next
	}

	fmt.Println("Progression:")
	for _, v := range progression {
		fmt.Println(v)
	}
}

// formatNumber converts a slice of integers into a formatted string.
func formatNumber(n []int) string {
	var parts []string
	for _, v := range n {
		parts = append(parts, fmt.Sprint(v))
	}
	return strings.Join(parts, ".")
}

// addNumbers adds two numbers represented as slices, with carry-over.
func addNumbers(a, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	result := make([]int, maxLen)
	carry := 0

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
