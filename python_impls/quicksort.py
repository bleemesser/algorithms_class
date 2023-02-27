import time
def quicksort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quicksort(left) + middle + quicksort(right)

start = time.time()
quicksort([5,2,7,9,4,2,6,1,6,9])
end = time.time()
print(f"Time taken: {end - start}")