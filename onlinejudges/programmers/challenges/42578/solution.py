# https://school.programmers.co.kr/learn/courses/30/lessons/42578


def solution(clothes):
    from collections import Counter

    counter = Counter([kind for _, kind in clothes])

    comb = 1
    for count in counter.values():
        comb *= count + 1
    return comb - 1
