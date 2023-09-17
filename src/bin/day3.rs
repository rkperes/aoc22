use std::collections::HashSet;
use std::io::read_to_string;
use std::io::BufRead;
use std::{char, io};

fn main() -> io::Result<()> {
    // problem1().unwrap();
    problem2().unwrap();

    Ok(())
}

fn problem2() -> io::Result<()> {
    let lines = read_to_string(io::stdin()).unwrap();
    let lines = lines.lines().collect::<Vec<_>>();

    let mut acc: u32 = 0;

    let group_size = 3;
    let groups = lines.chunks(group_size);

    for group in groups {
        let ch = letter_for_group(group);
        acc += ch_value(ch);
    }

    println!("Problem 2: {}", acc);

    Ok(())
}

fn letter_for_group(group: &[&str]) -> char {
    let mut previous: HashSet<char> = HashSet::new();
    let mut letters: HashSet<char> = HashSet::new();

    for (i, line) in group.iter().enumerate() {
        for ch in line.chars() {
            // in first line, no previous line to check
            if i == 0 {
                letters.insert(ch);
                continue;
            }

            // find letter that also was in all other groups
            if i == group.len() - 1 {
                if previous.contains(&ch) {
                    return ch;
                }

                continue;
            }

            // only keep item if it was already seen
            if previous.contains(&ch) {
                letters.insert(ch);
            }
        }

        previous = letters.clone();
        letters = HashSet::new();
    }

    ' '
}

fn problem1() -> io::Result<()> {
    let reader = io::BufReader::new(io::stdin());

    let mut acc: u32 = 0;

    for line in reader.lines() {
        let line = line.unwrap();
        let ch = find_common_letter_in_halfs(line.as_str());
        acc += ch_value(ch);
    }

    println!("Problem 1: {}", acc);

    Ok(())
}

fn find_common_letter_in_halfs(s: &str) -> char {
    let mut found: HashSet<char> = HashSet::new();

    for ch in s.chars().take(s.len() / 2) {
        found.insert(ch);
    }

    for ch in s.chars().skip(s.len() / 2) {
        if found.contains(&ch) {
            return ch;
        }
    }

    ' '
}

fn ch_value(ch: char) -> u32 {
    let chn = ch as u32;
    match ch {
        'a'..='z' => chn - ('a' as u32) + 1,
        'A'..='Z' => chn - ('A' as u32) + 26 + 1,
        _ => 0,
    }
}
