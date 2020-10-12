#!/usr/bin/env python3
# -*- encoding: utf-8 -*-

import sys


def counting_sort(data):
    max_value = max(data)
    count = [0 for _ in range(max_value + 1)]
    for digit in data:
        count[digit] += 1

    index = 0
    for value in range(max_value + 1):
        for _ in range(count[value]):
            data[index] = value
            index += 1

    return data


if __name__ == '__main__':
    N = int(sys.stdin.readline())
    data = [int(sys.stdin.readline()) for _ in range(N)]
    data = counting_sort(data)
    for item in data:
        print(item)
