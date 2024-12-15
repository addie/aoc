use std::{io};
use std::collections::HashSet;
use std::fs::File;
use std::io::BufRead;

const DAY: &str = "day10";
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

    part1(grid.clone());
    part2(grid.clone());

    Ok(())
}

fn part1(grid: Vec<Vec<i32>>) {
    let mut trailheads = Vec::new();
    for r in 0..grid.len() {
        for c in 0..grid[r].len() {
            if grid[r][c] == 0 {
                trailheads.push((r, c));
            }
        }
    }
    println!("trailheads- {:?}", trailheads);
    let mut sum_of_scores = 0;
    for trailhead in trailheads {
        let score = determine_score(&grid, trailhead);
        println!("score- {}", score);
        sum_of_scores += score;
    }
    println!("part1-{:?}", sum_of_scores);
}

fn part2(grid: Vec<Vec<i32>>) {
    let mut trailheads = Vec::new();
    for r in 0..grid.len() {
        for c in 0..grid[r].len() {
            if grid[r][c] == 0 {
                trailheads.push((r, c));
            }
        }
    }
    let mut sum_of_scores = 0;
    for trailhead in trailheads {
        let score = determine_score(&grid, trailhead);
        println!("score- {}", score);
        sum_of_scores += score;
    }
    println!("part1-{:?}", sum_of_scores);
}

fn determine_score(grid: &Vec<Vec<i32>>, trailhead: (usize, usize)) -> i32 {
    let start = trailhead;
    let mut stack = vec![start];
    let last_val = grid[start.0][start.1];
    let mut score = 0;
    let mut visited = HashSet::new();
    let dirs: Vec<(i32, i32)> = vec![(0, 1), (1, 0), (0, -1), (-1, 0)];
    while !stack.is_empty() {
        let curr = stack.pop().unwrap();
        let curr_value = grid[curr.0][curr.1];
        visited.insert(curr);
        if curr_value == 9 {
            score += 1;
            continue;
        }
        for dir in dirs.iter() {
            let next: (usize, usize) = ((curr.0 as i32 + dir.0) as usize, (curr.1 as i32 + dir.1) as usize);
            if is_valid(&grid, next) && !visited.contains(&next) && grid[next.0][next.1] == curr_value + 1 {
                stack.push(next);
            }
        }
    }

    score
}

fn is_valid(grid: &&Vec<Vec<i32>>, coord: (usize, usize)) -> bool {
    coord.0 < grid.len() && coord.1 < grid[0].len()
}