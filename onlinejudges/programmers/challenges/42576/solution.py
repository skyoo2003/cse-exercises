# https://school.programmers.co.kr/learn/courses/30/lessons/42576


def solution(participant, completion):
    race = {}
    for p in participant:
        if p not in race:
            race[p] = 0
        race[p] += 1

    finish = {}
    for c in completion:
        if c not in race:
            return c

        if c not in finish:
            finish[c] = 0
        finish[c] += 1

    for c in completion:
        if race[c] != finish[c]:
            return c

    for p in participant:
        if p not in finish:
            return p

    return ""
