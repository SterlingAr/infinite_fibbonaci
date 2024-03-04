package main

import (
	"fmt"
	"strings"
)

func main() {
    trigrams := map[string]int{
        "Heaven": 7, // 111
        "Lake":   5, // 101
        "Fire":   6, // 110
        "Thunder": 3, // 011
        "Wind":   4, // 100
        "Water":  2, // 010
        "Mountain": 1, // 001
        "Earth":  0, // 000
    }

    hexagrams := make(map[string]string)

    for lowerName, lowerValue := range trigrams {
        for upperName, upperValue := range trigrams {
            hexagramNumber := (upperValue << 3) | lowerValue
            hexagrams[fmt.Sprintf("%s over %s", upperName, lowerName)] = formatBinaryString(hexagramNumber)
        }
    }

    for name, binary := range hexagrams {
        fmt.Printf("%s: %s\n", name, binary)
    }
}

func formatBinaryString(n int) string {
    binary := fmt.Sprintf("%06b", n)
    return strings.ReplaceAll(binary, "0", "-")
}
