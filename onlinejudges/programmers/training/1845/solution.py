# https://school.programmers.co.kr/learn/courses/30/lessons/1845


def solution(nums):
    answer = 0

    pool = {}
    for n in nums:
        if n not in pool:
            pool[n] = 0
        pool[n] += 1

    unique = list(pool.keys())
    answer = len(unique) if len(unique) < len(nums) // 2 else len(nums) // 2

    return answer
