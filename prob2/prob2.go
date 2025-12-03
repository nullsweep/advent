package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func isInvalid(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)
	for patternLen := 1; patternLen <= len(str)/2; patternLen++ {
		if length%patternLen == 0 {
			pattern := str[0:patternLen]
			reps := length / patternLen
			repeat := strings.Repeat(pattern, reps)
			if repeat == str {
				return true
			}
		}
	}
	return false
}

func main() {
	sum := 0
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		ranges := strings.Split(line, ",")
		for _, rangeStr := range ranges {
			parts := strings.Split(rangeStr, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			for num := start; num <= end; num++ {
				if isInvalid(num) {
					sum += num
				}
			}
		}
	}
	log.Println(sum)
	return
}
