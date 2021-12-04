# Problem
#
# 有N(<=10) 个人，每个人水杯的容量是[v0, v1, ... , vN-1], (v[i] <= 10)
# 饮水机桶的容量为 V ( V < 100 and V >= max(v[i]) )
# 每个人打水时间是完全随机的。每次打水必须把自己水杯灌满
# 如果不够，就需要换桶水。如果正好用完，不需要换
# 初始状态，饮水机没有水
# 求第N个人第一次打水时需要换水的概率 ( 误差 < 1e-6 )

# Solve

import math

cups = [1, 2]
c0 = cups[-1]
N, V = len(cups), 100
M = 200

dp = [[[-1 for _ in range(V)] for _ in range(M)] for _ in range(N)]


def dfs(n: int, m: int, v: int) -> float:
    if dp[n][m][v] >= 0:
        return dp[n][m][v]
    if m == 0 and v == 0:
        return 1
    if m == 0 and v > 0:
        return 0
    if n == 0:
        if cups[n] * m != v:
            return 0
        return 1

    c = cups[n]
    res = 0
    for cm in range(m+1):
        dv = (v - (cm * c) % V + V) % V
        res += dfs(n-1, m - cm, dv) * math.comb(m, cm)
    dp[n][m][v] = res
    return res


def solve() -> float:
    res = 1/N
    for i in range(1, M):
        total, change = 0, 0
        for v in range(0, V):
            dv = V - v
            if dv == V:
                dv = 0
            cnt = dfs(N-1, i, v)
            if dv < c0:
                change += cnt
            total += cnt

        p = ((N-1)/N)**(i) * (1/N) * (change / total)
        res += p
    return res


print(solve())
