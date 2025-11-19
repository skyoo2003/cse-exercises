def sort(array: list[int]) -> list[int]:
    if len(array) <= 1:
        return array
    pivot = array[0]
    low, high = [], []
    for value in array[1:]:
        if value < pivot:
            low.append(value)
        else:
            high.append(value)
    return sort(low) + [pivot] + sort(high)


if __name__ == "__main__":
    import random

    array = []
    for _ in range(10):
        array.append(random.randrange(15))
    print("정렬 전: ", array)
    print("정렬 후: ", sort(array), flush=True)

