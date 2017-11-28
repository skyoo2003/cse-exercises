#!/usr/bin/env python3
# -*- encoding: utf-8 -*-

import sys
import heapq

def heap_sort(data):
    heapq.heapify(data)
    return [heapq.heappop(data) for _ in range(len(data))]

if __name__ == '__main__':
    N = int(sys.stdin.readline())
    data = [int(sys.stdin.readline()) for _ in range(N)]
    data = heap_sort(data)
    for item in data: print(item)
