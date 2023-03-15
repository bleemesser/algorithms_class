arr = [6, 5, 3, 1, 8, 7, 2, 4]

def insertion_sort(arr) -> list:
    for i in range(1, len(arr)):
        for j in range(i, 0, -1):
            if arr[j] < arr[j-1]:
                arr[j], arr[j-1] = arr[j-1], arr[j]
    return arr

print(insertion_sort(arr))