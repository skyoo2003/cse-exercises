# https://school.programmers.co.kr/learn/courses/30/lessons/42577


def solution(phone_book):
    answer = True

    phone_book = sorted(phone_book)
    for idx in range(len(phone_book) - 1):
        cb = phone_book[idx]
        nb = phone_book[idx + 1]
        if nb.startswith(cb):
            return False

    return answer
