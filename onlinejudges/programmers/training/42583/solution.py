# https://school.programmers.co.kr/learn/courses/30/lessons/42583


def solution(bridge_length, weight, truck_weights):
    from collections import deque

    elapsed_time = 0
    bridge = deque([0 for _ in range(bridge_length)], bridge_length)
    bridge_weight = 0
    current_truck = 0
    crossed_trucks = []
    while bridge:
        elapsed_time += 1

        crossed_truck_weight = bridge.popleft()
        bridge_weight -= crossed_truck_weight
        if crossed_truck_weight > 0:
            crossed_trucks.append(crossed_truck_weight)

        if current_truck >= len(truck_weights):
            continue

        truck_weight = truck_weights[current_truck]
        if bridge_weight + truck_weight <= weight:
            bridge.append(truck_weight)
            bridge_weight += truck_weight
            current_truck += 1
        else:
            bridge.append(0)

    return elapsed_time
