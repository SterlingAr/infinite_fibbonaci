package main

import (
	"fmt"
)

func main() {
    trigrams := []string{"111", "011", "101", "001", "110", "010", "100", "000"} // Fu Xi order

    for i, upper := range trigrams {
        for j, lower := range trigrams {
            hexagram := upper + lower
            hexagramNumber := i*8 + j + 1
            fmt.Printf("Hexagram %d: %s\n", hexagramNumber, hexagram)
        }
    }
}

