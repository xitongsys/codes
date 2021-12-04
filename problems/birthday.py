# Problem 1: 26 个人, 一年中每个月都有人过生日的概率
# 把N个人分到M个盒子里，每个盒子都有人的分法 
# 典型DP问题 dp[N][M]

N, M = 26, 12
F = [1] * (N + 1)
for i in range(1, N+1):
    F[i] = F[i-1] * i

def C(n:int, m:int):
    return F[n] // F[m] // F[n-m]

dp = [[0 for i in range(M+1)] for _ in range(N+1)]
for i in range(1, N+1):
    dp[i][1] = 1

for n in range(2, N+1):
    for m in range(2, M+1):
        for k in range(1, n-m+2):
            dp[n][m] += C(n, k) * dp[n-k][m-1]

total = 0
for m in range(1, M+1):
    total += C(M, m) * dp[N][m]

print(dp[N][M] / total) 
# 0.21547170566456322


#### brute force for check ######
total2, cnt2 = 0, 0
def dfs(bs:list):
    global total2, cnt2
    if len(bs) == N:
        total2 += 1
        if len(set(bs)) == M:
            cnt2 += 1
    else:
        for m in range(1, M+1):
            bs.append(m)
            dfs(bs)
            bs.pop()
dfs([])
print(cnt2 / total2)


# Problem 2: 考虑每月天数不同
# 类似前面，dp[n][m] 表示n个人分派到 mask 为 m 的月的概率
# mask m example: 分配 1，3月就是 101000000000

N, P = 26, [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31] 
P = [x / sum(P) for x in P]
M = len(P)
F = [1] * (N + 1)
for i in range(1, N+1):
    F[i] = F[i-1] * i

def C(n:int, m:int):
    return F[n] // F[m] // F[n-m]

dp = [[-1 for s in range(1<<M)] for _ in range(N+1)]
dp[0][0] = 1

def dfs(n:int, m:int)->float:
    if dp[n][m] >= 0:
        return dp[n][m]
    bm = bin(m).count('1')
    if n < bm:
        return 0
    res = 0
    for i in range(M):
        if (1<<i) & m:
            m2 = (1<<i) ^ m
            for a in range(1, n - bm + 2):
                res += C(n, a) * (P[i]**a) * dfs(n-a, m2)
            break
    dp[n][m] = res
    return res

print(dfs(N, (1<<M)-1))
# 0.21464194025884004


# Problem 3: 考虑闰年
# 四年一闰，百年不闰，四百年再闰
# 以400年为单位看，总天数146097，润天数97，普通天数146000
# 普通天出生概率P0 = 0.999336057550805   润天出生的概率为P1 = 0.000663942449194713
# 算法相同，只需要改下P数组即可

totalDay, normalDay, leapDay = 146097, 146000, 97
P0, P1 = normalDay / totalDay, leapDay / totalDay

# normal year
N, P = 26, [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31] 
P = [x / 365 * P0 for x in P]
# 0.2109672712321332

# leap year
P[1] = 28 / 365 * P0 + P1


M = len(P)
F = [1] * (N + 1)
for i in range(1, N+1):
    F[i] = F[i-1] * i

def C(n:int, m:int):
    return F[n] // F[m] // F[n-m]

dp = [[-1 for s in range(1<<M)] for _ in range(N+1)]
dp[0][0] = 1

def dfs(n:int, m:int)->float:
    if dp[n][m] >= 0:
        return dp[n][m]
    bm = bin(m).count('1')
    if n < bm:
        return 0
    res = 0
    for i in range(M):
        if (1<<i) & m:
            m2 = (1<<i) ^ m
            for a in range(1, n - bm + 2):
                res += C(n, a) * (P[i]**a) * dfs(n-a, m2)
            break
    dp[n][m] = res
    return res

print(dfs(N, (1<<M)-1))
# 0.21464194025884004
