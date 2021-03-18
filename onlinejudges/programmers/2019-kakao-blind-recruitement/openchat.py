def solution(record):
    answer = []

    utn = {}  # utn: user to nickname
    for item in record:
        cmd, uid, *rest = item.split()
        if cmd == 'Enter':
            utn[uid] = rest[0]
        elif cmd == 'Change':
            utn[uid] = rest[0]

    for item in record:
        cmd, uid, *rest = item.split()
        nickname = utn[uid]
        if cmd == 'Enter':
            answer.append(f'{nickname}님이 들어왔습니다.')
        elif cmd == 'Leave':
            answer.append(f'{nickname}님이 나갔습니다.')

    return answer
