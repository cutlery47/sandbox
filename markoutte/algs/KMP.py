def KMPSearch(pattern: str, text: str):
    pattern_length = len(pattern)
    pi = prefixFunction(pattern)
    q = 0
    for i in range(0, len(text) + 1):
        while q > 0 and pattern[q] != text[i]:
            q = pi[q]
            if q == None:
                print('123')

        if pattern[q] == text[i]:
            q += 1

        if q == pattern_length:
            return i - pattern_length + 1

    return -1


def prefixFunction(pattern: str) -> list:
    pattern_length = len(pattern)
    table = [0 for _ in range(pattern_length + 1)]
    k = 0
    for q in range(1, pattern_length):
        while k > 0 and pattern[k] != pattern[q]:
            k = table[k]

        if pattern[k] == pattern[q]:
            k += 1

        table[q + 1] = k

    return table