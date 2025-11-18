# https://school.programmers.co.kr/learn/courses/30/lessons/12909


def solution(s):
    answer = True

    st = []
    for c in s:
        if c == "(":
            st.append(c)
        else:
            try:
                st.pop()
            except IndexError:
                answer = False
                break

    if st:
        answer = False

    return answer
