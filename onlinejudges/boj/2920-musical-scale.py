#!/usr/bin/env python3

line = input()
items = [int(i) for i in line.strip().split(' ')]
status = ""

for i in range(0, len(items) - 1):
    diff = items[i + 1] - items[i]
    if diff > 0:
        if not status:
            status = "ascending"
        elif status != "ascending":
            status = "mixed"
    elif diff < 0:
        if not status:
            status = "descending"
        elif status != "descending":
            status = "mixed"

print(status)
