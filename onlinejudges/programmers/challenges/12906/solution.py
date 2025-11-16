# https://school.programmers.co.kr/learn/courses/30/lessons/12906?language=python3


def solution(arr):
    answer = []
    for value in arr:
        if answer and answer[-1] == value:
            continue
        answer.append(value)
    return answer
