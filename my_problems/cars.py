# Problem
# 一条单向公路上有N辆车，速度 vs = [v0, v1, ... ], 位置 ps = [p0, p1, ...]
# 若后面车追上前面车，则后面车变成与前面车相同的速度
# 求每辆车达到最终速度的时间


# Solve
# 如果某辆车速度比前面所有车速度都小，则这辆车是一个节点，它后面的车不会超过这辆车
# 按照节点，将车分组
# 在每组内，考虑第i，i+1车，如果 i 车达到最终速度时间为 t[i]，不考虑其他车，计算第i+1车追上本组第一辆车的时间 t'[i+1]

# 如果 t'[i+1] > t[i], 则 t[i+1] = t'[i+1]。因为如果i+1在达到最终速度前没有追上前车，那么时间就是t'[i+1]。如果i+1在达到
# 最终速度前已经追上了前车，那它达到最终速度的时间 = t[i], 而这个时间一定 > t'[i+1](因为它追上前车，速度一定下降了), 而这与
# t'[i+1] > t[i]矛盾。多以i+1一定在达到最终速度前没有追上前车

# 如果 t'[i+1] <= t[i], 则t[i+1] = t[i]。因为 i+1 车达到最终速度的时间t[i+1]一定 >= t[i], 而t'[i+1] <= t[i] 说明
# 在i+1达到最终速度之前，其一定已经追上了前一辆车
import sys
from typing import List


def solve(vs: List[int], ps: List[int]) -> List[float]:
    n = len(vs)
    nodes = []
    minv = 1000000
    for i in range(n):
        if vs[i] <= minv:
            minv = vs[i]
            nodes.append(i)

    res = [0] * n
    nn = len(nodes)
    for i in range(nn):
        b = nodes[i]
        e = n if i + 1 >= nn else nodes[i+1]
        for j in range(b+1, e):
            tj = (ps[j] - ps[b]) / (vs[j] - vs[b])
            if tj >= res[j-1]:
                res[j] = tj
            else:
                res[j] = res[j-1]
    return res

print(solve([1,2,3,4], [1,2,3,4]))
