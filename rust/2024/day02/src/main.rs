use std::fs::File;
use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    let file_path = "data.txt";
    let file = File::open(file_path)?;

    // Use a buffered reader
    let reader = io::BufReader::new(file);

    // total safe count
    let mut num_safe = 0;
    let mut num_safe_with_dampeners = 0;

    // Read the file line by line
    for line in reader.lines() {
        let line = line.unwrap();
        // Split the string by whitespace
        let mut levels: Vec<i32> = line
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();

        if is_safe(&levels) {
            num_safe += 1;
        }

        if is_safe_with_dampeners(&mut levels) {
            num_safe_with_dampeners += 1;
        }
    }

    println!("Safe Count: {}", num_safe);
    println!("Safe Count (with dampeners): {}", num_safe_with_dampeners);

    Ok(())
}

fn is_safe(levels: &Vec<i32>) -> bool {
    if levels[1] == levels[0] {
        return false;
    }
    if levels[1] - levels[0] > 0 { // increasing
        for i in 0..levels.len()-1 {
            let diff = levels[i + 1] - levels[i];
            if diff < 1 || diff > 3 {
                return false;
            }
        }
    } else if levels[1] - levels[0] < 0 { // decreasing
        for i in 0..levels.len()-1 {
            let diff = levels[i + 1] - levels[i];
            if diff > -1 || diff < -3 {
                return false;
            }
        }
    }
    true
}
fn is_safe_with_dampeners(levels: &Vec<i32>) -> bool {
    if is_safe(levels) {
        return true;
    }
    for i in 0..levels.len() {
        let dampened_levels: Vec<_> = levels[..i]
            .iter()
            .chain(levels[i + 1..].iter())
            .cloned()
            .collect();

        // alternative syntax
        // let dampened_levels: Vec<_> = levels
        //     .iter()
        //     .enumerate()
        //     .filter(|&(i, _)| i != i)
        //     .map(|(_, &value)| value)
        //     .collect();

        if is_safe(&dampened_levels) {
            return true;
        }
    }
    false
}
