use std::fs::File;
use std::io;

const DAY: &str = "day06";
pub fn solution() -> io::Result<()> {
    let file_path = format!("src/{}/data.txt", DAY);
    let file = File::open(file_path)?;

    // Use a buffered reader
    let reader = io::BufReader::new(file);

    Ok(())
}