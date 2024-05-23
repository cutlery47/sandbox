def prepare_matrix(matrix) -> None:
    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            matrix[i][j] /= 10

def determine_states(matrix) -> list:
    states = [num for num in range(1, len(matrix) + 1)]
    return states

