# Advent of Code 2021 - Go Solutions

My solutions for [Advent of Code 2021](https://adventofcode.com/2021) written in Go.

## Setup

Make sure you have Go installed (1.25+):

```bash
go version
```

## Running Solutions

Each day has a single main.go with both parts:

```bash
cd day01
go run main.go 1              # Part 1 with sample.txt (example data)
go run main.go 1 input.txt    # Part 1 with your personal puzzle input
go run main.go 2              # Part 2 with sample.txt
go run main.go 2 input.txt    # Part 2 with your personal puzzle input
```

## Input Files

Each day has two input files:
- `sample.txt` - Example data from the puzzle (safe to commit)
- `input.txt` - Your personal puzzle input (gitignored)

Develop and test with sample data, then run with your personal input for the final answer.

## Structure

```
day01/
├── main.go     # Both Part 1 and Part 2 solutions
├── sample.txt  # Example data from puzzle
└── input.txt   # Your personal puzzle input (gitignored)
day02/
├── main.go
├── sample.txt
└── input.txt
...
```