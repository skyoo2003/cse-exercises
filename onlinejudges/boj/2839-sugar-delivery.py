#!/usr/bin/env python3
# -*- encoding: utf-8 -*-

"""
2839-sugar-delivery.py
https://www.acmicpc.net/problem/2839
"""

N = int(input())

def sum(n5k, n3k):
    return 5 * n5k + 3 * n3k

n5k, n3k = 0, 0
succeed = False
while True:
    res = sum(n5k, n3k)
    if res == N:
        succeed = True
        break
    diff = N - res
    if diff % 5 == 0:
        n5k += 1
    elif diff % 3 == 0:
        n3k += 1
    elif diff >= 5:
        n5k += 1
    elif diff >= 3:
        n3k += 1
    else:
        break

print (n5k + n3k if succeed else -1)
