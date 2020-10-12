#!/usr/bin/env python3
# -*- encoding: utf-8 -*-

"""
1065-hansu.py
https://www.acmicpc.net/problem/1065
"""


def stairs(num):
    units = [int(i) for i in str(num)]
    diff = 0
    for i in range(len(units) - 1):
        ndiff = units[i + 1] - units[i]
        if i != 0 and diff != ndiff:
            return False
        diff = ndiff
    return True


N = int(input())
total = sum([1 for n in range(1, N + 1) if stairs(n)])
print(total)
