use aoc22::ioutil;
use core::fmt;
use regex::Regex;
use std::{
    error::Error,
    io::{self, Write},
};

fn main() -> io::Result<()> {
    problem1();

    Ok(())
}

fn problem1() {
    let lines = ioutil::stdin_as_lines();

    let mut acc: i32 = 0;

    for line in lines {}

    println!("Problem 1: {}", acc)
}

struct Stacks {
    element_matcher: Regex,
    s: Vec<Vec<char>>,
}

impl fmt::Display for Stacks {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for stack in self.s.into_iter() {
            for elem in stack {
                match write!(f, "[{}] ", elem) {
                    Ok(_) => (),
                    Err(e) => return Err(e),
                };
            }

            match write!(f, "{}", "\n") {
                Ok(_) => (),
                Err(e) => return Err(e),
            };
        }

        Ok(())
    }
}

impl Stacks {
    pub fn new() -> Self {
        Stacks {
            s: Vec::new(),
            element_matcher: Regex::new(r"^(?:(?:(\[\w\])|[\s]{3})\s)+").unwrap(),
        }
    }

    pub fn parse_line(line: String) -> Result<(), Box<dyn std::error::Error>> {
        Ok(())
    }
}
