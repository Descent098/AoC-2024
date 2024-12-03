from typing import List

def part1(data:str) -> int:
    left, right = parse_input_text(data)

    difference = 0
    for l,r in zip(left, right):
        difference += abs(int(l)-int(r))
    return difference
   
def part2(data:str) -> int:
    left, right = parse_input_text(data)

    similarity = 0
    for l in left:
        appearances = 0
        for r in right:
            if l == r:
                appearances+= 1
            if l <  r:
                break
        similarity += l * appearances
    return similarity 
    
def parse_input_text(data:str) ->(List[int], List[int]):
    left = []
    right = []
    for line in data.split("\n"):
        l, r = line.strip().split("   ")
        left.append(int(l))
        right.append(int(r))
    
    left.sort()
    right.sort()
    return left, right
    
if __name__ == "__main__":
    with open("input.txt", "r") as f:
        data = f.read()
        
    print(part1(data))
    print(part2(data))