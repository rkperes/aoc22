use std::io;
use std::io::BufRead;

fn main() -> io::Result<()> {
    // problem1().unwrap();
    problem2().unwrap();
    Ok(())
}

fn problem2() -> io::Result<()> {
    let reader = io::BufReader::new(io::stdin());

    let mut acc: u32 = 0;

    for line in reader.lines() {
        let line = line.unwrap();
        acc += points_from_match2(line);
    }

    println!("Problem 2: {}", acc);

    Ok(())
}

fn points_from_match2(input: String) -> u32 {
    let mut plays = input.split_whitespace();

    let theirs = play_from(plays.next().unwrap());
    let mine = decide_play(&theirs, plays.next().unwrap());

    mine as u32 + mine.score_vs(theirs)
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

    fn loses(&self) -> Play {
        match self {
            Play::Rock => Play::Paper,
            Play::Paper => Play::Scissors,
            Play::Scissors => Play::Rock,
        }
    }

    fn wins(&self) -> Play {
        match self {
            Play::Rock => Play::Scissors,
            Play::Paper => Play::Rock,
            Play::Scissors => Play::Paper,
        }
    }
}

fn decide_play(base: &Play, decision: &str) -> Play {
    match decision {
        "X" => base.wins(),
        "Y" => *base,
        "Z" => base.loses(),
        _ => panic!("bad play {}", decision),
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
