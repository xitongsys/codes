#[allow(unused_imports)]
use std::collections::BTreeMap;
use std::collections::BTreeSet;
use std::collections::HashMap;
use std::collections::HashSet;

#[derive(Default, Debug)]
struct Solution {}

impl Solution {
    fn solve(&mut self) {

    }

    fn main(&mut self) {
        let mut sc = io::Scanner::default();
        
    }
}

fn main() {
    let mut sol = Solution::default();
    sol.main();
}

mod io {
    use std::io::{stdin, stdout, BufWriter, Write};
    #[derive(Default)]
    pub struct Scanner {
        buffer: Vec<String>,
    }
    impl Scanner {
        pub fn next<T: std::str::FromStr>(&mut self) -> T {
            loop {
                if let Some(token) = self.buffer.pop() {
                    return token.parse().ok().expect("Failed parse");
                }
                let mut input = String::new();
                stdin().read_line(&mut input).expect("Failed read");
                self.buffer = input.split_whitespace().rev().map(String::from).collect();
            }
        }
    }
}
