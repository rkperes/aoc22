use std::io;
use std::io::BufRead;

fn main() -> io::Result<()> {
    // problem1().unwrap();
    problem2().unwrap();
    Ok(())
}

fn problem2() -> io::Result<()> {
    let reader = io::BufReader::new(io::stdin());

    let top3: &mut [i32] = &mut [0; 3];

    let mut acc: i32 = 0;

    for line in reader.lines() {
        match line?.trim().parse::<i32>() {
            Ok(val) => acc += val,
            Err(_) => {
                update_top_n(top3, acc);
                acc = 0;
            }
        };
    }
    update_top_n(top3, acc);

    println!("Problem 2: {}", top3[0] + top3[1] + top3[2]);
    Ok(())
}

fn update_top_n(top_n: &mut [i32], val: i32) {
    let mut next = val;

    for i in 0..top_n.len() {
        let cur = *top_n.get(i).unwrap();
        if next > cur {
            let tmp: i32;
            (tmp, next) = (next, cur);
            top_n[i] = tmp;
        }
    }
}

fn problem1() -> io::Result<()> {
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
