# https://school.programmers.co.kr/learn/courses/30/lessons/42586


def solution(progresses, speeds):
    answer = []

    ets = []
    for progress, speed in zip(progresses, speeds):
        et = (100 - progress) / speed
        if et % 1 != 0:
            et = round(et + 0.5)
        et = int(et)
        ets.append(et)

    c_et = ets[0]
    cnt = 0
    for et in ets:
        if et <= c_et:
            cnt += 1
        else:
            answer.append(cnt)
            cnt = 1
            c_et = et
    answer.append(cnt)

    return answer
