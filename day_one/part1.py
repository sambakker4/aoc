def get_lists(filename):
    list1 = []
    list2 = []

    with open(filename, "r") as f:
        lines = f.readlines()

        for line in lines:
            num1, num2 = line.split()
            num1, num2 = int(num1), int(num2)
            list1.append(num1)
            list2.append(num2)
    return list1, list2

if __name__ == "part1.py":
    list1, list2 = get_lists("input.txt")
    list1.sort()
    list2.sort()

    distances = []

    for i in range(len(list1)):
        distances.append(abs(list1[i] - list2[i]))

    total_distance = sum(distances)
    print(total_distance)