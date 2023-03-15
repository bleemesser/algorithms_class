import time

def countingsort(arr):
    mi, ma = min(arr), max(arr)
    buckets = [0] * (ma - mi + 1)
    for i in arr:
        buckets[i - mi] += 1
    # use only 1 for loop to return the sorted array
    out = []
    for i in range(len(buckets)):
        out += [i + mi] * buckets[i]
    return out
    
    # return [i + mi for i in range(len(buckets)) for _ in range(buckets[i])]

print(countingsort([5,2,7,9,4,2,6,1,6,9]))