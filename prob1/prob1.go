package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	cp := 50
	count := 0
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		if direction == 'R' {
			cp = cp + distance
		} else if direction == 'L' {
			cp = cp - distance
		}

		cp = cp % 100
		if cp < 0 {
			cp = cp + 100
		}
		if cp == 0 {
			count++
		}
	}
	log.Println("Times landed on 0:", count)
}
