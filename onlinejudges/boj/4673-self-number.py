#!/usr/bin/env python3
# -*- encoding: utf-8 -*-

"""
4673-self-number.py
https://www.acmicpc.net/problem/4673
"""


def d(num):
    total = num
    for u in str(num):
        total += int(u)
    return total


MAX_NUM = 10000

flag = [False] * (MAX_NUM + 1)
for i in range(MAX_NUM):
    num = i + 1
    dnum = d(num)
    if dnum <= MAX_NUM:
        flag[dnum] = True

for i in range(MAX_NUM):
    num = i + 1
    if not flag[num]:
        print(num)
