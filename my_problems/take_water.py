# Problem
#
# 有N(<=10) 个人，每个人水杯的容量是[v0, v1, ... , vN-1], (v[i] <= 10)
# 饮水机桶的容量为 V ( V < 100 and V >= max(v[i]) )
# 每个人打水时间是完全随机的, 每次打水必须把自己水杯灌满
# 如果不够，就需要换桶水。如果正好用完，不需要换
# 初始状态，饮水机没有水
# 求第N个人第一次打水时需要换水的概率 ( 误差 < 1e-6 )

# Solve
# 因为打水时间完全随机，所有人按照打水次序排成一个序列，0,1,0,2,1,2,0,.....
# 第N个人第一次出现，可能在0，1，2，。。。。。个位置
# 第N个人第一次出现在 i 位置的概率为 ((N-1)/N)**(i) * (1/N)
# 而第N个人出现在 i 位置时候，前面已经喝过的总水量有多种不同的情形。而最终影响第N个人的是 (总水量 % V) 的值
# dp[n][m][v] 表示 到 index = n 个人，从其中选 m 个人排在第N个人前面，这时候 (喝水总量 % V) == v 的情况个数

import math

cups = [1, 2, 4, 5, 3, 4, 10, 2]
c0 = cups[-1]
N, V = len(cups), 20
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
        if (cups[n] * m) % V != v:
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
    res = 0
    for i in range(0, M):
        total, change = 0, 0
        for v in range(0, V):
            dv = V - v
            if dv == V:
                dv = 0
            cnt = dfs(N-2, i, v)
            if dv < c0:
                change += cnt
            total += cnt

        p = ((N-1)/N)**(i) * (1/N) * (change / total)
        res += p
    return res


print(solve())

#### check special case ( only 2 people ) ####


def check():
    res = 0
    c0, c1 = cups[0], cups[1]
    for i in range(0, M):
        v = V - (c0 * i) % V
        if v == V:
            v = 0
        if v < c1:
            res += ((N-1)/N) ** i * (1/N)
    return res


print(check())
