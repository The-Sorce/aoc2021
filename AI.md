# AI Assistant Instructions

This is an Advent of Code 2021 repository with Go solutions.

## Project Structure
- Each day has its own directory (`day01/`, `day02/`, etc.)
- Each day contains `main.go`, `sample.txt`, and `input.txt`
- Simple, flat structure - no over-engineering

## Input Files
- `sample.txt` - Example data from puzzle (safe to commit)
- `input.txt` - Personal puzzle input (gitignored, never commit)

## When helping:
- Keep Go code simple and beginner-friendly
- Don't solve puzzles completely - provide guidance and skeleton code
- Focus on Go basics: file reading, string parsing, loops, slices
- Encourage testing with sample data first, then personal input
- Use `go fmt` to format code when making changes

## Commands:
- `go run main.go 1` - Run Part 1 with sample.txt
- `go run main.go 1 input.txt` - Run Part 1 with personal input
- `go run main.go 2` - Run Part 2 with sample.txt
- `go run main.go 2 input.txt` - Run Part 2 with personal input
- `go fmt` - Format Go code
- `go mod tidy` - Clean up module dependencies

## Code Structure
Each main.go contains:
- `main()` function with argument parsing and shared input reading
- `solvePart1(numbers []int)` function for Part 1 logic
- `solvePart2(numbers []int)` function for Part 2 logic