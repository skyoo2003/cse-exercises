def sort(array: list[int]) -> list[int]:
    for i in range(len(array) - 1):
        min_i = i
        for j in range(i + 1, len(array)):
            if array[j] < array[min_i]:
                min_i = j
        array[i], array[min_i] = array[min_i], array[i]
    return array


if __name__ == "__main__":
    import random

    array = []
    for _ in range(10):
        array.append(random.randrange(15))
    print("정렬 전: ", array)
    print("정렬 후: ", sort(array), flush=True)

