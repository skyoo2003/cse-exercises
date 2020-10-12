where = int(input())

left = where
curr = 1

while left > curr:
    left -= curr
    curr += 1

num, denom = 0, 0  # 분자, 분모
is_even = curr % 2
if is_even:
    num, denom = (curr + 1) - left, left
else:
    num, denom = left, (curr + 1) - left

print("{}/{}".format(num, denom))
