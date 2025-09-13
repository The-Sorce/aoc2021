# AI Assistant Instructions

This is an Advent of Code 2021 repository with both Go and Python solutions for learning both languages.

## Project Structure
- Each day has its own directory (`day01/`, `day02/`, etc.)
- Each day contains `main.go`, `main.py`, `sample.txt`, and `input.txt`
- Both solutions implement identical logic for easy comparison
- Simple, flat structure - no over-engineering

## Input Files
- `sample.txt` - Example data from puzzle (safe to commit)
- `input.txt` - Personal puzzle input (gitignored, never commit)

## When helping:
- Keep both Go and Python code simple and beginner-friendly
- Don't solve puzzles completely - provide guidance and skeleton code
- Focus on language basics: file reading, string parsing, loops, data structures
- Encourage testing with sample data first, then personal input
- Use `go fmt` for Go code and follow PEP 8 for Python
- When creating solutions, maintain identical logic between both languages

## Commands:

### Go:
- `go run main.go 1` - Run Part 1 with sample.txt
- `go run main.go 1 input.txt` - Run Part 1 with personal input
- `go run main.go 2` - Run Part 2 with sample.txt
- `go run main.go 2 input.txt` - Run Part 2 with personal input
- `go fmt` - Format Go code
- `go mod tidy` - Clean up module dependencies

### Python:
- `python3 main.py 1` - Run Part 1 with sample.txt
- `python3 main.py 1 input.txt` - Run Part 1 with personal input
- `python3 main.py 2` - Run Part 2 with sample.txt
- `python3 main.py 2 input.txt` - Run Part 2 with personal input

## Code Structure

### Go (main.go):
- `main()` function with argument parsing and shared input reading
- `solvePart1(numbers []int)` function for Part 1 logic
- `solvePart2(numbers []int)` function for Part 2 logic

### Python (main.py):
- `main()` function with argument parsing and shared input reading
- `solve_part1(numbers)` function for Part 1 logic
- `solve_part2(numbers)` function for Part 2 logic
- `if __name__ == "__main__":` guard for script execution