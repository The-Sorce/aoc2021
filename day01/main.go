package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <1|2> [filename]")
		return
	}

	part := os.Args[1]
	filename := "sample.txt"
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Printf("Input data length: %d lines\n", len(lines))

	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	switch part {
	case "1":
		solvePart1(numbers)
	case "2":
		solvePart2(numbers)
	default:
		fmt.Println("Invalid part. Use '1' or '2'")
	}
}

func solvePart1(numbers []int) {
	increases := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			increases++
		}
	}
	fmt.Println("Answer:", increases)
}

func solvePart2(numbers []int) {
	slidingWindowSums := make([]int, len(numbers)-2)
	for i := 3; i <= len(numbers); i++ {
		slidingWindowSums[i-3] = numbers[i-3] + numbers[i-2] + numbers[i-1]
	}
	solvePart1(slidingWindowSums)
}
