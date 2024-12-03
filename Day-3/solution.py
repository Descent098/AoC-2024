import re

def part1(data:str) -> int:
    result = 0
    regex = r"mul\((\d{1,3})\,(\d{1,3})\)"

    matches = re.finditer(regex, data, re.MULTILINE)
    for _, match in enumerate(matches, start=1):
        left, right = match.group()[4:-1].split(',')
        result = result + (int(left) * int(right))
    
    return result
        

def part2(data:str) -> int:
    result = 0
    regex = r"mul\((\d{1,3})\,(\d{1,3})\)|do\(\)|don\'t\(\)"

    matches = re.finditer(regex, data, re.MULTILINE)
    do = True

    for _, match in enumerate(matches, start=1):
        if match.group() == "do()":
            do = True
            continue
        elif match.group() == "don't()":
            do = False
            continue
        if do:
            left, right = match.group()[4:-1].split(',')
            result = result + (int(left) * int(right))

    return result
        

if __name__ == "__main__":
    
    with open("input.txt") as f:
        temp = f.read()
        
        
    print(f"Part 1: {part1(temp)}")
    print(f"Part 2: {part2(temp)}")
