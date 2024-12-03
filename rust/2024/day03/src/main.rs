use std::fs::File;
use std::io::{self, BufRead};
use regex::Regex;

fn main() -> io::Result<()> {
    let file_path = "data.txt";
    let file = File::open(file_path)?;

    // Use a buffered reader
    let reader = io::BufReader::new(file);

    let mut sum_of_mults = 0;
    let mut sum_of_mults_2 = 0;

    // Read the file line by line
    let mut mult_enabled = true;
    for line in reader.lines() {
        let line = line.unwrap();

        // part1
        let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
        for captures in re.captures_iter(&line) {
            if let (Some(first), Some(second)) = (captures.get(1), captures.get(2)) {
                sum_of_mults += first.as_str().parse::<i32>().unwrap() * second.as_str().parse::<i32>().unwrap();
            }
        }

        // part2

        // This works by creating 3 regex to capture mult(x,y), do(), and don't()
        // as well as a combined regex that matches on any of those three
        // then it iterates the string looking for a match on the combined regex.
        // If it finds a match, it will check which of the 3 regex hit
        // and adds the logic where we enable/disable mult processing.
        let regex_a = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
        let regex_b = Regex::new(r"do\(\)").unwrap();
        let regex_c = Regex::new(r"don't\(\)").unwrap();
        let combined_pattern = Regex::new(r"mul\((\d+),(\d+)\)|do\(\)|don't\(\)").unwrap();

        for mat in combined_pattern.find_iter(&line) {
            let matched_text = mat.as_str();
            if mult_enabled && regex_a.is_match(matched_text) {
                // println!("Matched 'mult' pattern: {}", matched_text);
                if let Some(captures) = regex_a.captures(matched_text) {
                    if let (Some(first), Some(second)) = (captures.get(1), captures.get(2)) {
                        let one = first.as_str().parse::<i32>().unwrap();
                        let two = second.as_str().parse::<i32>().unwrap();
                        sum_of_mults_2 += one * two;
                    }
                }
            } else if regex_b.is_match(matched_text) {
                mult_enabled = true;
            } else if regex_c.is_match(matched_text) {
                mult_enabled = false;
            }
        }
    }
    println!("Part 1: {}", sum_of_mults);
    println!("Part 2: {}", sum_of_mults_2);

    Ok(())
}

