from queue import Queue


def is_valid_index(i: int, j: int, m: int, n: int) -> bool:
    return 0 <= i < m and 0 <= j < n


def destory(
    q: Queue[tuple[int, int]],
    matrix: list[list[int]],
):
    m = len(matrix)
    n = len(matrix[0])
    while not q.empty():
        idx = q.get()

        if matrix[idx[0]][idx[1]] == 1:
            # 把 1 翻转为 0
            # 把周围 8 个位置的 1 也翻转为 0
            si = idx[0] - 1
            sj = idx[1] - 1
            for i in range(3):
                for j in range(3):
                    ni = si + i
                    nj = sj + j
                    if is_valid_index(ni, nj, m, n):
                        v = matrix[ni][nj]
                        if v == 1:
                            print(f"({ni}, {nj}) 置为0")
                            matrix[ni][nj] = 0
                            for x in range(3):
                                for y in range(3):
                                    if is_valid_index(ni + x, nj + y, m, n):
                                        q.put((ni + x, nj + y))


def run(matrix: list[list[int]]) -> int:
    ans = 0
    q = Queue()

    i = 0
    while i < len(matrix[0]):
        j = 0
        while j < len(matrix):
            if matrix[i][j] == 1:
                q.put((i, j))
                ans += 1
                destory(q, matrix)
                print(f"====== 第 {ans} 轮结束 ======")
            j += 1
        i += 1
    return ans


if __name__ == "__main__":
    """
    1 0 0 1 0
    1 1 1 0 0
    0 0 0 1 0
    0 0 0 0 1
    1 0 0 0 0
    """
    m = [
        [1, 0, 0, 1, 0],
        [1, 1, 1, 0, 0],
        [0, 0, 0, 1, 0],
        [0, 0, 0, 0, 1],
        [1, 0, 0, 0, 0],
    ]
    assert run(m) == 2
