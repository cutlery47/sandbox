import random
from typing import Optional

def chain(matrix, states, steps, start: Optional[int] = None):
    result = [random.choice(seq=states) if not start else start]

    state = result[0]
    for i in range(steps):
        state = random.choices(population=states, weights=matrix[state - 1], k=1)[0]
        result.append(state)

    return result
