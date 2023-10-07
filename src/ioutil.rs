use std::io;

pub fn stdin_as_lines() -> Vec<String> {
    let lines = io::read_to_string(io::stdin()).unwrap();
    lines.lines().map(String::from).collect()
}
