from part1 import get_lists

left_list, right_list = get_lists("input.txt")

total_score = 0

for num in left_list:
    score = 0
    for num2 in right_list:
        if num2 == num:
            score += 1
    total_score += (score * num)
print(total_score)