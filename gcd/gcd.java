    static long gcd(long a, long b){
        if(a == 0) return b;
        return gcd(b % a, a);
    }

    static long exGCD(long a, long b) {
        long x = 0, y = 1, u = 1, v = 0;
        long gcd = 0;
        while(a != 0){
            long q = b / a, r = b % a;
            long m = x - u * q, n = y - v * q;
            b = a;
            a = r;
            x = u;
            y = v;
            u = m;
            v = n;
            gcd = b;
        }
        return x;
    }

    static long inv(long a, long mod){
        long x = exGCD(a, mod);
        x = (x % mod + mod) % mod;
        return x;
    }
