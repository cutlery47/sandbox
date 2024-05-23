from tvims.utils import prepare_matrix, determine_states
from tvims.excercise_6 import chain
from tvims.excercise_7 import percentage

if __name__ == "__main__":
    matrix = [
        [7, 0, 1, 0, 2, 0, 0, 0],
        [0, 2, 0, 0, 0, 3, 4, 1],
        [2, 0, 2, 2, 4, 0, 0, 0],
        [0, 2, 0, 5, 0, 0, 0, 3],
        [1, 0, 4, 2, 2, 1, 0, 0],
        [0, 2, 0, 3, 0, 5, 0, 0],
        [0, 6, 0, 0, 0, 1, 2, 1],
        [0, 2, 0, 5, 0, 0, 1, 2]
    ]

    prepare_matrix(matrix)
    states = determine_states(matrix)

    for steps in [10, 50, 100, 1000]:
        print(chain(matrix=matrix, states=states, steps=steps))

    percentage(matrix=matrix, states=states, steps=100_000_000)
