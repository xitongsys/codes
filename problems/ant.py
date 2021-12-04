# Problem 1
# 长度为L棍子上有N只蚂蚁，位置[p0,p1,p2,...], 速度[v0,v1,v2...]，速度可正可负，但不为零
# 两只蚂蚁相遇后，交换速度
# 多久棍子上没有蚂蚁

# Solve: 交换速度 = 穿越
from typing import List


def solve(L: int, ps: List[int], vs: List[int]):
    n = len(ps)
    res = 0
    for i in range(n):
        p, v = ps[i], vs[i]
        l = p if v < 0 else L - p
        res = max(res, l / abs(v))
    return res


print(solve(10, [1, 2, 3], [1, -1, 2]))

# Problem 2
# 求每只蚂蚁掉落的时间
# Solve:
# 蚂蚁无法穿越，所以一定先从两边顺序掉落
# 所有掉落时间和掉落左右可以确定，分别分配给左右蚂蚁即可


def solve(L: int, ps: List[int], vs: List[int]):
    n = len(ps)
    ts = []
    for i in range(n):
        p, v = ps[i], vs[i]
        l = p if v < 0 else L - p
        ts.append(l / v)
    ts.sort(key=lambda x: abs(x))
    res = [0] * n
    i, j = 0, n - 1
    for t in ts:
        if t < 0:
            res[i] = abs(t)
            i += 1
        else:
            res[j] = t
            j -= 1
    return res


print(solve(10, [1, 2, 3], [1, -1, 2]))
