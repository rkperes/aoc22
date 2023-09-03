use std::io;
use std::io::BufRead;

fn main() -> io::Result<()> {
    problem1().unwrap();
    // problem2().unwrap();
    Ok(())
}

fn problem1() -> io::Result<()> {
    let reader = io::BufReader::new(io::stdin());

    let mut acc: u32 = 0;

    for line in reader.lines() {
        let line = line.unwrap();
        acc += points_from_match(line);
    }

    println!("Problem 1: {}", acc);

    Ok(())
}

fn points_from_match(input: String) -> u32 {
    let mut plays = input.split_whitespace();

    let theirs = play_from(plays.next().unwrap());
    let mine = play_from(plays.next().unwrap());

    mine as u32 + mine.score_vs(theirs)
}

#[derive(Copy, Clone)]
enum Play {
    Rock = 1,
    Paper,
    Scissors,
}

impl Play {
    fn score_vs(&self, other: Play) -> u32 {
        match (self, other) {
            (Play::Rock, Play::Scissors)
            | (Play::Paper, Play::Rock)
            | (Play::Scissors, Play::Paper) => 6,

            (Play::Rock, Play::Rock)
            | (Play::Paper, Play::Paper)
            | (Play::Scissors, Play::Scissors) => 3,

            _ => 0,
        }
    }
}

fn play_from(s: &str) -> Play {
    match s {
        "A" | "X" => Play::Rock,
        "B" | "Y" => Play::Paper,
        "C" | "Z" => Play::Scissors,

        _ => panic!("bad play {}", s),
    }
}
