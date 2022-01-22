# NxM的方格，任取不同4个点，连成正方形的个数
# 1. N < 1e6, M < 1e6
# 2. N, M 无限制 [2, 1e18]

# 只考虑一个 n x n 正方形边上的点，所能连成的正方形个数为 (n-1)
# 只需要统计各个边长正方形的个数，分别乘以（n-1）再求和即可
# 不防设 N > M
# count = 求和 (N-x)*(M-x)*(x+1), 其中，0 <= x <= M 的整数
# 此式子有解析解


def cal(N, M):
    if M > N:
        N, M = M, N
    res = 0
    for x in range(M):
        res += (N-x) * (M-x) * (x+1)
    return res