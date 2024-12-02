from part1 import get_list, safe_report

if __name__ == "__main__":
    lists = get_list("input.txt")
    safe = 0
    for lst in lists:

        for i in range(len(lst)):
            tmp = lst.pop(i)
            if safe_report(lst):
                safe += 1
                lst.insert(i, tmp)
                break
            lst.insert(i, tmp)
    print(safe)

