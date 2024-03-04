package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const maxGoroutines = 13
const tzolkin = 260

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		// Randomly decide the number of concurrent goroutines
		numGoroutines := rand.Intn(maxGoroutines) + 1

		outputChannel := make(chan string, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go generateRandomIChingSequence(outputChannel)
		}

		// Continuously read from the output channel for a random duration
		timeout := time.After(time.Duration(rand.Intn(tzolkin)) * time.Millisecond)
		for {
			select {
			case output := <-outputChannel:
				fmt.Print(output)
				time.Sleep(time.Duration(rand.Intn(tzolkin)) * time.Millisecond) // Random pause
			case <-timeout:
				break
			}
		}
		close(outputChannel) // Close the current batch of goroutines
	}
}

func generateRandomIChingSequence(outputChannel chan<- string) {
	yinYang := []string{"1", "0"}

	for {
		var sequenceLength int

		// Randomly decide the type of sequence
		sequenceType := rand.Intn(3) // 0 (season), 1 (trigram), or 2 (hexagram)

		switch sequenceType {
		case 0:
			sequenceLength = 2
		case 1:
			sequenceLength = 3
		case 2:
			sequenceLength = 6
		}

		var builder strings.Builder
		for i := 0; i < sequenceLength; i++ {
			builder.WriteString(yinYang[rand.Intn(2)])
		}

		outputChannel <- builder.String() + " "

		// Random delay before generating the next sequence
		time.Sleep(time.Duration(rand.Intn(tzolkin) + 1) * time.Millisecond)
	}
}
