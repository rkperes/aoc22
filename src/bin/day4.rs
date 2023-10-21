use aoc22::ioutil;
use core::fmt;
use std::io::{self, Write};

fn main() -> io::Result<()> {
    problem2();

    Ok(())
}

fn problem1() {
    let lines = ioutil::stdin_as_lines();

    let mut acc: i32 = 0;

    for line in lines {
        let p = Pair::from_string(line);
        if p.contains_self() {
            acc += 1;
        }
    }

    println!("Problem 1: {}", acc)
}

fn problem2() {
    let lines = ioutil::stdin_as_lines();

    let mut acc: i32 = 0;

    for line in lines {
        let p = Pair::from_string(line);
        if p.overlaps_self() {
            // println!("overlaps {}", p);
            acc += 1;
        } else {
            // println!("not overlaps {}", p);
        }
    }

    println!("Problem 2: {}", acc)
}

struct Pair {
    first: Range,
    second: Range,
}

impl fmt::Display for Pair {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}, {}", self.first, self.second)
    }
}

impl Pair {
    fn new(first: Range, second: Range) -> Self {
        Self { first, second }
    }

    fn from_string(s: String) -> Self {
        let parts = s.split_once(',').unwrap();
        Self::new(
            Range::from_string(String::from(parts.0)),
            Range::from_string(String::from(parts.1)),
        )
    }

    pub fn contains_self(&self) -> bool {
        self.first.contains(&self.second) || self.second.contains(&self.first)
    }

    pub fn overlaps_self(&self) -> bool {
        self.first.overlaps(&self.second)
    }
}

struct Range(i32, i32);

impl fmt::Display for Range {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}-{}", self.0, self.1)
    }
}

impl Range {
    pub fn new(start: i32, end: i32) -> Self {
        Self(start, end)
    }

    pub fn from_string(s: String) -> Self {
        let parts = s.split_once('-').unwrap();
        Self::new(
            parts.0.parse::<i32>().unwrap(),
            parts.1.parse::<i32>().unwrap(),
        )
    }

    pub fn contains(&self, other: &Range) -> bool {
        self.0 <= other.0 && self.1 >= other.1
    }

    pub fn overlaps(&self, other: &Range) -> bool {
        (self.1 >= other.0 && self.1 <= other.1) || (other.1 >= self.0 && other.1 <= self.1)
    }
}
