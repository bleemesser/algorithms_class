def is_match(input, pattern):
    n = len(input)
    print("n:", n)
    m = len(pattern)
    print("m:", m)

    T = [[False] * (m+1) for _ in range(n+1)]

    T[0][0] = True

    for j in range(1, m+1):
        if pattern[j-1] == '*':
            T[0][j] = T[0][j-1]

    for i in range(1, n+1):
        for j in range(1, m+1):
            if pattern[j-1] == '*':
                T[i][j] = T[i-1][j] or T[i][j-1]
            elif pattern[j-1] == '?' or pattern[j-1] == input[i-1]:
                T[i][j] = T[i-1][j-1]

    # print the dp table
    print("DP table:")
    for row in T:
        print(' '.join(map(str, row)))

    # find the start and end indices of the pattern in the input string
    found_pattern_start = -1
    found_pattern_end = -1

    pattern_indices = []
    for i in range(n+1):
        for j in range(m+1):
            if T[i][j] and i <= m and i not in pattern_indices:
                pattern_indices.append(i)
        # print the pattern indices
    print("Pattern indices:", pattern_indices)


    if len(pattern_indices) == 0:
        return False, found_pattern_start, found_pattern_end


    found_pattern_start = pattern_indices[0]
    found_pattern_end = pattern_indices[-1] - 1

    return True, found_pattern_start, found_pattern_end

input = "xzyb"
pattern = "zy"

b, start, end = is_match(input, pattern)
print(b)
for i in range(start, end+1):
    print(input[i], end='')
print()
