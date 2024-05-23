import random
from excercise_6 import chain

def percentage(matrix, states, steps):
    state_hits = {
        1: 0,
        2: 0,
        3: 0,
        4: 0,
        5: 0,
        6: 0,
        7: 0,
        8: 0
    }
    huge_chain = chain(matrix=matrix, states=states, steps=steps)

    for el in huge_chain:
        state_hits[el] += 1

    for key in state_hits.keys():
        print(f"The probability to hit state {key} is {state_hits[key] / len(huge_chain)}")



