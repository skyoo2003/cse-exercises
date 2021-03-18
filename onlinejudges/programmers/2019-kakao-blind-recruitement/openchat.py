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

if __name__ == '__main__':
    inp = ["Enter uid1234 Muzi", "Enter uid4567 Prodo","Leave uid1234","Enter uid1234 Prodo","Change uid4567 Ryan"]
    exp_res = ["Prodo님이 들어왔습니다.", "Ryan님이 들어왔습니다.", "Prodo님이 나갔습니다.", "Prodo님이 들어왔습니다."]
    act_res = solution(inp)
    assert exp_res == act_res, 'Invalid result'

