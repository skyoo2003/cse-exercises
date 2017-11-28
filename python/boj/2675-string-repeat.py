N = int(input())

while N > 0:
    (rep, word) = tuple(input().split()[:2])
    rep = int(rep)
    out = []
    for ch in word:
        out.append(ch * rep)
    print ("".join(out))
    N -= 1
