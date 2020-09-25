T = int(input())
while T > 0:
    (H, W, N) = tuple(int(i) for i in input().split()[:3])
    print ("%d%02d" % (N%H if N%H != 0 else H, int(N/H)+1 if N%H != 0 else int(N/H)))
    T -= 1
