import re
def get_text(filename):
    with open(filename, "r") as f:
        lines = f.read()
        return lines

def find_sets(lines):
    
    expression = "mul\(([0-9]{1,3}),([0-9]{1,3})\)"

    return re.findall(expression, lines)


            
            

if __name__ == "__main__":
    text = get_text("input.txt")
    sets = find_sets(text)
    total = 0
    for num1, num2 in sets:
        total += (int(num1) * int(num2))
    print(total)