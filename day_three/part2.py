from part1 import get_text, find_sets
import re
if __name__ == "__main__":
    expression = r"(do(?:n't)?\(\))|(mul\([0-9]{1,3},[0-9]{1,3}\))"
    text = get_text("input.txt")
    matches = re.findall(expression, text)
    total = 0
    do = True

    for match in matches:
        if match[0] == "":
            if do:
                nums = re.findall("mul\(([0-9]{1,3}),([0-9]{1,3})\)", match[1])
                total += int(nums[0][0]) * int(nums[0][1])
        else:
            if match[0] == "do()":
                do = True
            if match[0] == "don't()":
                do = False
    print(total)