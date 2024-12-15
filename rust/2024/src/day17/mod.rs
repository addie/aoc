use std::{io};
use std::fs::File;
use std::io::BufRead;

const DAY: &str = "day17";
pub fn solution() -> io::Result<()> {
    let file_path = format!("src/{}/data_sample.txt", DAY);
    // let file_path = format!("src/{}/data.txt", DAY);
    let file = File::open(file_path)?;

    // Use a buffered reader
    let reader = io::BufReader::new(file);

    // Create graph
    let mut grid: Vec<Vec<i32>> = Vec::new();
    for line in reader.lines() {
        let line = line?;
        let row = line.chars().map(|c| c.to_digit(10).unwrap() as i32).collect::<Vec<i32>>();
        grid.push(row);
    }

    // part1(grid.clone());
    // part2(grid.clone());

    Ok(())
}
