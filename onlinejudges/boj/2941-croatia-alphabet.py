line = input()
transforms = ["c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z=", ]
cnt = 0

while line:
    if line[:3] in transforms:
        line = line[3:]
    elif line[:2] in transforms:
        line = line[2:]
    else:
        line = line[1:]
    cnt += 1

print(cnt)
