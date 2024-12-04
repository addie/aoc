use std::fs::File;
use std::io::{self, BufRead};

pub fn solution() -> io::Result<()> {
    let day = "day04";

    let file_path = format!("src/{}/data.txt", day);
    let file = File::open(file_path)?;

    // Use a buffered reader
    let reader = io::BufReader::new(file);

    // Read the file line by line
    for line in reader.lines() {
        let line = line?;

    }
    // println!("Day04 Part1: {}", _);
    // println!("Day04 Part2: {}", _);
    Ok(())
}
