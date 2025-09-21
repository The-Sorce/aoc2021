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

	switch part {
	case "1":
		solvePart1(lines)
	case "2":
		solvePart2(lines)
	default:
		fmt.Println("Invalid part. Use '1' or '2'")
	}
}

func solvePart1(lines []string) {
	pos := 0
	depth := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		direction := fields[0]
		units, _ := strconv.Atoi(fields[1])

		switch direction {
		case "forward":
			pos += units
		case "down":
			depth += units
		case "up":
			depth -= units
		}
	}
	fmt.Println("Answer:", pos*depth)
}

func solvePart2(lines []string) {
	pos := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		direction := fields[0]
		units, _ := strconv.Atoi(fields[1])

		switch direction {
		case "forward":
			pos += units
			depth += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	fmt.Println("Answer:", pos*depth)
}
