use std::io;
use std::io::BufRead;

fn main() -> io::Result<()> {
    let reader = io::BufReader::new(io::stdin());

    let mut acc: i32 = 0;
    let mut max: i32 = 0;

    for line in reader.lines() {
        match line?.trim().parse::<i32>() {
            Ok(val) => acc += val,
            Err(_) => {
                if acc > max {
                    max = acc;
                }
                acc = 0;
                continue;
            }
        };
    }

    if acc > max {
        max = acc;
    }

    println!("max: {}", max);
    Ok(())
}
