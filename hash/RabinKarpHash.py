# LeetCode1044
# 给你一个字符串 s ，考虑其所有 重复子串 ：即 s 的（连续）子串，在 s 中出现 2 次或更多次。这些出现之间可能存在重叠。
# 返回 任意一个 可能具有最长长度的重复子串。如果 s 不含重复子串，那么答案为 "" 。
#
# RobinKarp Hash
# hash_value = sum(S[i] * C^(N-i-1)),  C is a const


class Solution:
    def longestDupSubstring(self, s: str) -> str:
        MOD = int(1e16 + 7)
        def mpow(a, n):
            if n == 0:
                return 1
            r = mpow(a, n//2)
            r = r * r
            if n % 2 == 1:
                r *= a
            return r % MOD

        C = 261
        n = len(s)
        l, r = 1, n - 1
        while l <= r:
            m = l + (r-l)//2
            st = set()
            
            # RobinKarp
            h = 0
            for i in range(m):
                h = h * C + ord(s[i])
                h %= MOD
            st.add(h)
            f = False
            for i in range(1, n+1-m):
                h = (h - ord(s[i-1]) * mpow(C, m-1)) * C + ord(s[i+m-1])   
                h %= MOD 
                if h in st:
                    f = True
                    break
                st.add(h)
            if f:
                l = m + 1
            else:
                r = m - 1

        if r == 0:
            return ""
        
        # RobinKarp
        h, st = 0, set()
        for i in range(r):
            h = h * C + ord(s[i])
            h %= MOD
        st.add(h)

        for i in range(1, n+1-r):
            h = (h - ord(s[i-1]) * mpow(C, r-1)) * C + ord(s[i+r-1])    
            h %= MOD
            if h in st:
                return s[i:i+r]
            st.add(h)
        return ""
