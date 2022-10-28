use std::{
    io::{self, Read},
    iter::Scan,
};

struct Scanner {
    it: std::vec::IntoIter<String>,
}

impl Scanner {
    fn new() -> Scanner {
        let mut buf = String::new();
        std::io::stdin().read_to_string(&mut buf);

        return Scanner {
            it: {
                std::io::stdin().read_to_string(&mut buf).unwrap();
                buf.split_whitespace().map(|x|x.to_string()).collect::<Vec<String>>().into_iter()
            }
        };
    }

    fn next<T: std::str::FromStr>(&mut self) -> T 
    where
    <T as std::str::FromStr>::Err: std::fmt::Debug,{
        return self.it.next().unwrap().parse().unwrap()
    }
}

#[derive(Default, Debug)]
struct Solution {
}

impl Solution {
    fn solve(&mut self) {
    }
    fn main(&mut self) {
        let mut sc = Scanner::new();
    }
}

fn main() {
    let mut sol = Solution::default();
    sol.main();
}
