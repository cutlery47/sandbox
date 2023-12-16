def naiveSearch(pattern: str, text: str) -> int:
    for i in range(len(text) - len(pattern) + 1):
        if equals(pattern, text, i, i + len(pattern)):
            return i

    return -1

def equals(pattern: str, text: str, start: int, finish: int) -> bool:
    for i in range(start, finish):
        if pattern[i - start] != text[i]:
            return False

    return True