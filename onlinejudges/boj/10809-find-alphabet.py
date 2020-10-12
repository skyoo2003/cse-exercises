import string

alphabets = string.ascii_lowercase[:26]
indices = {key: -1 for key in alphabets}

line = input()
for idx, ch in enumerate(line):
    if indices[ch] == -1:
        indices[ch] = idx

print(' '.join([str(indices[ch]) for ch in alphabets]))
