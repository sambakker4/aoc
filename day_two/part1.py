def get_list(filename):
    with open(filename, "r") as f:
        lines = f.readlines()
        lst = []
        for line in lines:
            lst2 = []
            for char in line.split():
                lst2.append(int(char))
            lst.append(lst2)
        return lst

def safe_report(lst):
    inc = False
    dec = False
    for i in range(1, len(lst)):

        diff = lst[i] - lst[i - 1]
        
        if 1 <= diff <= 3:
            inc = True
        elif -3 <= diff <= -1:
            dec = True
        else:
            break

        if dec and inc:
            break
    else:
        return True
    return False

if __name__ == "__main__":
    lists = get_list("input.txt")
    safe = 0

    for lst in lists:
        if safe_report(lst):
            safe += 1
    print(safe)