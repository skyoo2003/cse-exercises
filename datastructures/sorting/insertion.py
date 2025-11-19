def sort(array: list[int]) -> list[int]:
    # METHOD 1: swap
    # for i in range(1, len(array)):
    #     while i > 0 and array[i] < array[i - 1]:
    #         array[i], array[i - 1] = array[i - 1], array[i]
    #         i -= 1
    #
    # METHOD 2: shift and assign
    for i in range(1, len(array)):
        value = array[i]
        while i > 0 and value < array[i - 1]:
            array[i] = array[i - 1]
            i -= 1
        array[i] = value
    return array


if __name__ == "__main__":
    import random

    array = []
    for _ in range(10):
        array.append(random.randrange(15))
    print("정렬 전: ", array)
    print("정렬 후: ", sort(array), flush=True)
