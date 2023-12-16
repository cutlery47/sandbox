def BMSearch(pattern: str, text: str) -> int:
    text_length = len(text)
    pattern_length = len(pattern)
    table = dict()
    for i in range(len(pattern) - 1):
        table[pattern[i]] = i

        for i in range(text_length - pattern_length + 1):
            for j in range(pattern_length - 1, -1, -1):
                if ord(pattern[j]) != ord(text[i + j]):
                    move = table.get(text[i + j], None)
                    i += max(1, j - (pattern_length if not move else move))
                    continue
            return i

        return -1