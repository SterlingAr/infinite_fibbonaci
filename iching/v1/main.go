package main

import (
	"fmt"
)

func main() {
    yinYang := map[string]string{
        "M": "♂",
        "F":  "♀",
    }

    visualRepresentation := map[string]string{
        "1": "M",
        "0": "F",
    }

  // Generate the four seasons
  for name1, binary1 := range yinYang {
	for name2, binary2 := range yinYang {
		seasonName := name1 + name2
		seasonSymbol := visualRepresentation[binary1] + visualRepresentation[binary2]
		seasonBinary := binary1 + binary2
		fmt.Printf("Season %s: %s (%s)\n", seasonName, seasonSymbol, seasonBinary)
	}
}
fmt.Println()

// Generate the eight trigrams
for season1, binary1 := range yinYang {
	for season2, binary2 := range yinYang {
		for season3, binary3 := range yinYang {
			trigramName := season1 + season2 + season3
			trigramSymbol := visualRepresentation[binary1] + visualRepresentation[binary2] + visualRepresentation[binary3]
			trigramBinary := binary1 + binary2 + binary3
			fmt.Printf("Trigram %s: %s (%s)\n", trigramName, trigramSymbol, trigramBinary)
		}
	}
}
fmt.Println()

// Generate the sixty-four hexagrams
for trigram1 := 0; trigram1 < 8; trigram1++ {
	for trigram2 := 0; trigram2 < 8; trigram2++ {
		hexagramSymbol := rune(0x4DC0 + trigram1*8 + trigram2)
		hexagramBinary := fmt.Sprintf("%03b%03b", trigram1, trigram2)
		fmt.Printf("Hexagram: %s (%s)\n", string(hexagramSymbol), hexagramBinary)
	}
}}

