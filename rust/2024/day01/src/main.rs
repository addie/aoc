use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    let file_path = "data.txt";
    let file = File::open(file_path)?;

    // Use a buffered reader
    let reader = io::BufReader::new(file);

    let mut ids_one: Vec<i32> = Vec::new();
    let mut ids_two: Vec<i32> = Vec::new();

    // Read the file line by line
    for line in reader.lines() {
        let line = line.unwrap();
        // Split the string by whitespace
        let str_ids: Vec<&str> = line.split_whitespace().collect();
        // Convert &str to an integer
        ids_one.push(str_ids[0].parse().unwrap());
        ids_two.push(str_ids[1].parse().unwrap());
    }
    ids_one.sort();
    ids_two.sort();

    let total: i32 = ids_one
        .iter()
        .zip(ids_two.iter())
        .map(|(&id1, &id2)| (id1 - id2).abs())
        .sum();
    println!("total part 1 = {}", total);

    // Part 2
    // Create a counter using HashMap
    let mut total2: i32 = 0;
    let mut counter: HashMap<i32, i32> = HashMap::new();
    for &num in &ids_two {
        *counter.entry(num).or_insert(0) += 1;
    }
    for &id in &ids_one {
        if let Some(&count) = counter.get(&id) {
            println!("ID: {}, Count: {}", id, count);
            total2 += id * count;
        }
    }
    println!("total part 2 = {}", total2);

    Ok(())
}
