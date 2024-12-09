with open("input1.txt", "r") as f:
    input_lines = f.readlines()

lines = [l.replace("\n", "").split("   ") for l in input_lines]

print(lines)
