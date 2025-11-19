# NOT WORKING YET...


def merge(array: list[int], left: int, center: int, right: int):
    l, r = 0, 0
    l_array = array[left:center]
    r_array = array[center : right + 1]
    while l < len(l_array) and r < len(r_array):
        if array[l] <= array[r]:
            array[left] = l_array[l]
            l += 1
        else:
            array[left] = r_array[r]
            r += 1
        left += 1

    while l < len(l_array):
        array[left] = l_array[l]
        l += 1
        left += 1

    while r < len(r_array):
        array[left] = r_array[r]
        r += 1
        left += 1


def do(array: list[int], left: int, right: int):
    if left == right:
        return
    center = (left + right) // 2
    do(array, left, center)
    do(array, center + 1, right)
    merge(array, left, center + 1, right)


def sort(array: list[int]) -> list[int]:
    do(array, 0, len(array) - 1)
    return array


if __name__ == "__main__":
    import random

    array = []
    for _ in range(10):
        array.append(random.randrange(15))
    print("정렬 전: ", array)
    print("정렬 후: ", sort(array), flush=True)
