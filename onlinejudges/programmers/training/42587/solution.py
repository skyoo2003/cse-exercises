# https://school.programmers.co.kr/learn/courses/30/lessons/42587


def solution(priorities, location):
    from collections import deque

    count = 0
    queue = deque([(p, i) for i, p in enumerate(priorities)])
    while queue:
        priority, idx = queue.popleft()

        if any(p > priority for p, _ in queue):
            queue.append((priority, idx))
        else:
            count += 1
            if idx == location:
                return count

    return count
