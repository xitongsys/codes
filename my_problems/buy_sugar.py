# 有 N 块糖果
# 每块糖果每斤价钱为 P[p0, p1, ... , pN]
# 你对每块糖果的打分为 V[v0, v1, ..., vN]
# 现在给你 F 元钱, 要求你买 Q 斤糖果，使得所买糖果 总得分 最大
# 请你返回每种糖果所买的斤数 q[q0, q1, ..., qN]
#
# 即 sum(qi) = Q, sum(qi * pi) <= F 
# 求 max(sum(vi * qi))
#
# 数据范围
# N ~ [1, 1000]
# pi ~ float (0, 1e9)
# vi ~ float (0, 1e9)
# Q ~ float (0, 1e9)
# F ~ float (0, 1e9)

# solve
#
# 假如买了3种糖果分别为 q1, q2, q3
# 对应 p1 < p2 < p3 
# 我们可以减少1，3，增加2，同样满足 sum(pi * qi) = F
# 而这样操作后，总得分 要么增加，要么减少
# 如果增加，那么就一直这样操作，直到1,3中一个减为0
# 如果减少，就反向操作，2->1,3 直到 q2 = 0
# 这样操作总会使 总喜欢程度增加, 而且买糖果的种类从3种变成了2种
# 以此类推，总得分最大，只有两种情况 买 1 种糖果 或者 买 2 种糖果
# 枚举比较即可

from typing import List
def solve(P:List[float], V:List[float], F:float, Q:float):
    N = len(P)
    mxv, res = 0, [0] * N
    for i in range(N):
        if Q * P[i] > F:
            continue
        if Q * V[i] > mxv:
            mxv = Q * V[i]
            res = [0] * N
            res[i] = Q
    for i in range(N):
        for j in range(i+1, N):
            qi = (P[i] * Q - F) / (P[i] - P[j])
            qj = (F - P[j]*Q) / (P[i] - P[j])
            v = qi * V[i] + qj * V[j]
            if qi >= 0 and qj >= 0 and v > mxv:
                mxv = v
                res = [0] * N
                res[i], res[j] = qi, qj
    return res


