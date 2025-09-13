# Advent of Code 2021 - Go & Python Solutions

My solutions for [Advent of Code 2021](https://adventofcode.com/2021) written in both Go and Python for learning both languages.

## Setup

Make sure you have Go installed (1.25+):

```bash
go version
```

And Python 3.11+:

```bash
python3 --version
```

## Running Solutions

Each day has both main.go and main.py with identical functionality:

### Go
```bash
cd day01
go run main.go 1              # Part 1 with sample.txt (example data)
go run main.go 1 input.txt    # Part 1 with your personal puzzle input
go run main.go 2              # Part 2 with sample.txt
go run main.go 2 input.txt    # Part 2 with your personal puzzle input
```

### Python
```bash
cd day01
python3 main.py 1             # Part 1 with sample.txt (example data)
python3 main.py 1 input.txt   # Part 1 with your personal puzzle input
python3 main.py 2             # Part 2 with sample.txt
python3 main.py 2 input.txt   # Part 2 with your personal puzzle input
```

## Input Files

Each day has two input files:
- `sample.txt` - Example data from the puzzle (safe to commit)
- `input.txt` - Your personal puzzle input (gitignored)

Develop and test with sample data, then run with your personal input for the final answer.

## Structure

```
day01/
├── main.go     # Go solutions (Part 1 and Part 2)
├── main.py     # Python solutions (Part 1 and Part 2)
├── sample.txt  # Example data from puzzle (shared)
└── input.txt   # Your personal puzzle input (gitignored)
day02/
├── main.go
├── main.py
├── sample.txt
└── input.txt
...
```