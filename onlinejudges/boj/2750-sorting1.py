#!/usr/bin/env python3
# -*- encoding: utf-8 -*-
import sys


def quick_sort(data=[]):
    MAX_LEVELS = len(data)
    rdata = [data[i] for i in range(MAX_LEVELS)]
    begin = [0 for i in range(MAX_LEVELS)]
    end = [MAX_LEVELS for i in range(MAX_LEVELS)]
    i = 0
    while i >= 0:
        l = begin[i]
        r = end[i] - 1
        if l < r:
            pivot = rdata[l]
            while l < r:
                while rdata[r] >= pivot and l < r:
                    r = r - 1
                if l < r:
                    rdata[l] = rdata[r]
                    l = l + 1
                while rdata[l] <= pivot and l < r:
                    l = l + 1
                if l < r:
                    rdata[r] = rdata[l]
                    r = r - 1
            rdata[l] = pivot
            begin[i + 1] = l + 1
            end[i + 1] = end[i]
            end[i] = l
            i = i + 1
            if end[i] - begin[i] > end[i - 1] - begin[i - 1]:
                begin[i], begin[i - 1] = begin[i - 1], begin[i]
                end[i], end[i - 1] = end[i - 1], end[i]
        else:
            i = i - 1
    return rdata


if __name__ == '__main__':
    N = int(sys.stdin.readline())
    data = [int(sys.stdin.readline()) for _ in range(N)]
    data = quick_sort(data)
    for item in data:
        print(item)
