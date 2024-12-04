use std::fs::File;
use std::io::{self, BufRead};

pub fn solution() -> io::Result<()> {
    let day = "day04";

    let file_path = format!("src/{}/data.txt", day);
    let file = File::open(file_path)?;

    // Use a buffered reader
    let reader = io::BufReader::new(file);

    // Create graph
    let mut grid: Vec<Vec<char>> = Vec::new();
    for line in reader.lines() {
        let line = line?;
        let row = line.chars().collect::<Vec<char>>();
        grid.push(row);
    }

    let count1 = part1(&grid);
    let count2 = part2(&grid);

    println!("Day04 Part1: {}", count1);
    println!("Day04 Part2: {}", count2);
    Ok(())
}

fn part1(grid: &Vec<Vec<char>>) -> i32 {
    const DIRS: [(i32, i32); 8] = [
        (0, 1),
        (0, -1),
        (1, 0),
        (-1, 0),
        (-1, -1),
        (1, 1),
        (1, -1),
        (-1, 1),
    ];

    const TARGET: &str = "XMAS";

    let mut count = 0;
    for (r, row) in grid.iter().enumerate() {
        for (c, &col) in row.iter().enumerate() {
            if grid[r][c] != 'X' {
                continue;
            }
            // if we see an X search for XMAS
            // println!("Found X at ({}, {}), start DFS", r, c);
            for (dr, dc) in &DIRS {
                // println!("checking {},{}", dr, dc);
                let mut stack = Vec::new();
                stack.push((r as i32, c as i32, 0));
                while let Some((r, c, len)) = stack.pop() {
                    if len == TARGET.len() - 1 {
                        count += 1;
                        // println!("found XMAS, count {}", count);
                        break;
                    }
                    let (nr, nc) = (r + dr, c + dc);
                    // println!("checking grid pos ({},{}) for target char {}", nr, nc, TARGET.chars().nth(len+1).unwrap());
                    if in_bound(nr, nc, &grid)
                        && grid[nr as usize][nc as usize] == TARGET.chars().nth(len + 1).unwrap()
                    {
                        // println!("found target char {} at grid pos ({},{})", TARGET.chars().nth(len+1).unwrap(), nr, nc);
                        stack.push((nr, nc, len + 1));
                    }
                }
            }
        }
    }

    count
}

fn part2(grid: &Vec<Vec<char>>) -> i32 {
    let mut count = 0;
    for (r, row) in grid.iter().enumerate() {
        for (c, &col) in row.iter().enumerate() {
            if grid[r][c] != 'A' || r < 1 || c < 1 || r > grid.len() - 2 || c > grid[r].len() - 2 {
                continue;
            }
            if is_valid_x_mas(grid, r as i32, c as i32) {
                count += 1;
            }
        }
    }

    count
}

fn is_valid_x_mas(grid: &[Vec<char>], r: i32, c: i32) -> bool {
    let diag1 = (grid[(r - 1) as usize][(c - 1) as usize], grid[(r + 1) as usize][(c + 1) as usize]);
    let diag2 = (grid[(r - 1) as usize][(c + 1) as usize], grid[(r + 1) as usize][(c - 1) as usize]);

    (diag1 == ('M', 'S') || diag1 == ('S', 'M')) && (diag2 == ('M', 'S') || diag2 == ('S', 'M'))
}

fn in_bound(r: i32, c: i32, grid: &Vec<Vec<char>>) -> bool {
    r >= 0 && c >= 0 && r < grid.len() as i32 && c < grid[0].len() as i32
}
