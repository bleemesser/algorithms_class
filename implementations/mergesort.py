import time


def mergesort(arr):
    if len(arr) <= 1:
        return arr
    mid = len(arr) // 2
    left = mergesort(arr[:mid])
    right = mergesort(arr[mid:])
    li = 0
    ri = 0
    result = []
    while li < len(left) and ri < len(right):
        if left[li] <= right[ri]:
            result.append(left[li])
            li += 1
        else:
            result.append(right[ri])
            ri += 1
    result += left[li:]
    result += right[ri:]
    return result


start = time.time()
mergesort([5,2,7,9,4,2,6,1,6,9])
end = time.time()
print(f"Time taken: {end - start} seconds")


