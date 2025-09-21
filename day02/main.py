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
    
    with open(filename, 'r') as f:
        lines = f.read().strip().split('\n')
    
    print(f"Input data length: {len(lines)} lines")
    
    if part == "1":
        solve_part1(lines)
    elif part == "2":
        solve_part2(lines)
    else:
        print("Invalid part. Use '1' or '2'")

def solve_part1(lines):
    pos = 0
    depth = 0

    for line in lines:
        parts = line.split()
        direction, units = parts[0], int(parts[1])

        match direction:
            case "forward":
                pos += units
            case "down":
                depth += units
            case "up":
                depth -= units

    print("Answer:", pos*depth)

def solve_part2(lines):
    pos = 0
    depth = 0
    aim = 0

    for line in lines:
        parts = line.split()
        direction, units = parts[0], int(parts[1])

        match direction:
            case "forward":
                pos += units
                depth += aim * units
            case "down":
                aim += units
            case "up":
                aim -= units

    print("Answer:", pos*depth)

if __name__ == "__main__":
    main()