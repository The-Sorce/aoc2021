#!/usr/bin/env python3
import sys

def main():
    if len(sys.argv) < 2:
        print("Usage: python3 main.py <1|2> [filename]")
        return
    
    part = sys.argv[1]
    filename = "sample.txt"
    if len(sys.argv) > 2:
        filename = sys.argv[2]
    
    try:
        with open(filename, 'r') as f:
            lines = f.read().strip().split('\n')
    except Exception as e:
        print(f"Error reading file {filename}: {e}")
        return
    
    print(f"Input data length: {len(lines)} lines")
    
    try:
        numbers = [int(line) for line in lines]
    except ValueError as e:
        print(f"Error parsing numbers: {e}")
        return
    
    if part == "1":
        solve_part1(numbers)
    elif part == "2":
        solve_part2(numbers)
    else:
        print("Invalid part. Use '1' or '2'")

def solve_part1(numbers):
    increases = 0
    for i in range(1, len(numbers)):
        if numbers[i] > numbers[i-1]:
            increases += 1
    print("Answer:", increases)

def solve_part2(numbers):
    sliding_window_sums = []
    for i in range(len(numbers) - 2):
        sliding_window_sums.append(numbers[i] + numbers[i+1] + numbers[i+2])
    solve_part1(sliding_window_sums)

if __name__ == "__main__":
    main()