package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	totalSum := 0
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		max := 0
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				firstDigit := int(line[i] - '0')
				secondDigit := int(line[j] - '0')
				jolt := firstDigit*10 + secondDigit
				if jolt > max {
					max = jolt
				}
			}
		}
		totalSum += max
	}
	println(totalSum)
}
