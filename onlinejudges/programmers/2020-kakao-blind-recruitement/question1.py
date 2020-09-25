#!/usr/bin/env python3
# https://programmers.co.kr/learn/courses/30/lessons/60057

def solution(s):
    answer = 0

    N = len(s)
    fz = None
    for l in range(1, N//2+1):
        ts = [s[i:i+l] for i in range(0, N, l)]
        lt = ts[0] # last term
        cnt = 1
        z = ''
        for t in ts[1:]:
            if lt == t:
                cnt += 1
            else:
                if cnt > 1:
                    z += str(cnt)
                z += lt
                lt = t
                cnt = 1
        if cnt > 1:
            z += str(cnt)
        z += lt
        if not fz or len(z) < len(fz):
            fz = z
    answer = len(fz)
    return answer


if __name__ == '__main__':
    s = 'aabbaccc'
    r = solution(s)
    print(r, r == 7)
    s = 'ababcdcdababcdcd'
    r = solution(s)
    print(r, r == 9)
    s = 'abcabcdede'
    r = solution(s)
    print(r, r == 8)
    s = 'abcabcabcabcdededededede'
    r = solution(s)
    print(r, r == 14)
    s = 'xababcdcdababcdcd'
    r = solution(s)
    print(r, r == 17)
