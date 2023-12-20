from naive import equals

def pseudoRollingHash(string: str, start: int, finish: int) -> int:
    result = 0
    for i in range(start, finish):
        result += ord(string[i])
    return result

def RKSearch(pattern: str, text: str) -> int:
    pattern_length = len(pattern)
    pattern_hash = pseudoRollingHash(pattern, 0, pattern_length)
    text_length = len(text)
    text_hash = pseudoRollingHash(text, 0, pattern_length)

    for i in range(text_length - pattern_length + 1):
        if pattern_hash == text_hash and equals(pattern, text, i, i + pattern_length):
            return i

        text_hash -= ord(text[i])
        text_hash += ord(text[i + pattern_length])

    return -1

