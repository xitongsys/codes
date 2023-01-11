# 给定 N（N <= 18) 个数的数组 A，用 这 N 个数排成一个长度为 M 的数列 B
# 要求 B 中任意区间中所有数字出现的频率至少有一个是奇数
# 例如 
# [1,2,3,2,1] 是一个符合条件的序列
# 而 [1,2,3,2,3] 不是一个符合条件的序列，比如区间[1:5] = [2,3,2,3]
# 其中 2，3 都出现了偶数次

# 如果无法生成这样的序列，返回空数组

# 数据范围：

# N ~ [2,18]
# M ~ [1, 10^6]


# Solve:
# 用一个 N 位的二进制数表示各个数出现次数的奇偶性
# M 个这样的二进制数组成一个数组 C，i 处的数表示从开始到 i 每个数字出现的奇偶性
# 任意区间中每个数字出现频率至少有一个是奇数，那么就要求 C 中每个数字都不一样
# 并且相邻两个数字只能由一位不同，which is famous known as Grey Code
# 因此最长可以是 (1<<N)-1

def grey_code(N:int)->list:
    res = [0, 1]
    for i in range(2, N+1):
        n = len(res)
        for j in range(n-1, -1, -1):
            res.append(res[j] ^ (1<<i))
    return res

def solve(A:list, M:int)->list:
    N = len(A)
    if M > (1<<N) - 1:
        return []
    gcs = grey_code(N)
    res = []
    for i in range(1, len(gcs)):
        m = gcs[i] ^ gcs[i-1]
        for j in range(0, N):
            if m & (1<<j) > 0:
                res.append(A[j])
                break
    return res



