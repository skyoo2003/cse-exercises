(A, B) = tuple([int(i[::-1]) for i in input().split()[:2]])
print(A if A > B else B)
