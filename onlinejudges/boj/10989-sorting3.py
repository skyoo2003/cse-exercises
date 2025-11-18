#!/usr/bin/env python3
# -*- encoding: utf-8 -*-

import sys


if __name__ == "__main__":
    N = int(sys.stdin.readline())
    data = (int(sys.stdin.readline()) for _ in range(N))
    counter = [0] * 10000
    for value in data:
        counter[value - 1] += 1

    for i in range(len(counter)):
        count = counter[i]
        if counter == 0:
            continue
        for _ in range(count):
            print(i + 1)
