scores = [int(input()) for i in range(5)]
scores = [i if i >= 40 else 40 for i in scores]
avg = int(sum(scores) / len(scores))

print(avg)
