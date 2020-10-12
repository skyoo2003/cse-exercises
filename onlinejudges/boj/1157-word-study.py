counters = {}

line = input()
for ch in line:
    ch = ch.lower()
    if ch not in counters:
        counters[ch] = 0
    counters[ch] += 1

(mch, mcnt, dup) = "", -1, False
for ch, cnt in counters.items():
    if cnt > mcnt:
        mch, mcnt, dup = ch, cnt, False
    elif cnt == mcnt:
        dup = True

print(mch.upper() if not dup else "?")
