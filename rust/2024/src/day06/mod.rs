use std::fs::File;
use std::io::{self, BufRead};

const DAY: &str = "day06";

#[derive(Clone, Copy, PartialEq, Debug)]
enum Dir {
    North,
    East,
    South,
    West,
}

impl Dir {
    fn next(&self) -> Self {
        match self {
            &Self::North => Self::East,
            &Self::East => Self::South,
            &Self::South => Self::West,
            &Self::West => Self::North,
        }
    }

    fn offset(&self) -> (i32, i32) {
        match self {
            Dir::North => (-1, 0),
            Dir::East => (0, 1),
            Dir::South => (1, 0),
            Dir::West => (0, -1),
        }
    }
}

pub fn solution() -> io::Result<()> {
    let file_path = format!("src/{}/data.txt", DAY);
    let file = File::open(file_path)?;

    // Use a buffered reader
    let reader = io::BufReader::new(file);
    let mut grid: Vec<Vec<char>> = Vec::new();
    for line in reader.lines() {
        let line = line?;
        let row = line.chars().collect::<Vec<char>>();
        grid.push(row);
    }

    // Find start position
    let mut count = 1; // make sure we count start position
    let mut curr: (usize, usize) = (0, 0);
    let mut facing = Dir::North;
    for r in 0..grid.len() {
        for c in 0..grid[r].len() {
            if grid[r][c] == '^' {
                curr = (r, c);
                grid[r][c] = 'X';
                // println!("visited ({}, {})", r, c);
                // pretty_print(&grid);
            }
        }
    }
    curr = ((curr.0 as i32 + Dir::North.offset().0) as usize, (curr.1 as i32 + Dir::North.offset().1) as usize);
    let mut next: (usize, usize) = curr;
    while in_bound(&grid, curr.0, curr.1) {
        if grid[curr.0][curr.1] == '.' {
            count += 1;
            grid[curr.0][curr.1] = 'X';
        }
        // println!("curr ({}, {})", curr.0, curr.1);
        curr = next;
        // println!("new curr ({}, {})", curr.0, curr.1);

        next = get_next(&mut curr, &mut facing);
        if !in_bound(&grid, next.0, next.1) {
            if grid[curr.0][curr.1] == '.' {
                count += 1;
            }
            break;
        }
        if grid[next.0][next.1] == '#' {
            facing = facing.next();
            next = get_next(&mut curr, &mut facing);
            // // println!("face change {:?}", facing)
        }
        // println!("new next ({}, {})", next.0, next.1);
        // pretty_print(&grid);
    }

    println!("{}", count);

    Ok(())
}

fn get_next(curr: &mut (usize, usize), moving: &mut Dir) -> (usize, usize) {
    ((curr.0 as i32 + moving.offset().0) as usize, (curr.1 as i32 + moving.offset().1) as usize)
}

fn in_bound(grid: &Vec<Vec<char>>, r: usize, c: usize) -> bool {
    r < grid.len() && c < grid[r].len()
}

fn pretty_print<T: std::fmt::Display>(matrix: &[Vec<T>]) {
    println!("  0 1 2 3 4 5 6 7 8 9");
    for (i, row) in matrix.iter().enumerate() {
        print!("{} ", i);
        // Join the elements of the row into a single string, separated by spaces.
        let line = row.iter()
            .map(|elem| elem.to_string())
            .collect::<Vec<_>>()
            .join(" ");
        println!("{}", line);
    }
    println!();
}