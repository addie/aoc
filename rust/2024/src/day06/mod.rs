use std::collections::{HashMap, HashSet};
use std::{fs, io};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
enum Direction {
    North,
    East,
    South,
    West,
}

impl Direction {
    fn turn_right(&self) -> Self {
        match self {
            Direction::North => Direction::East,
            Direction::East => Direction::South,
            Direction::South => Direction::West,
            Direction::West => Direction::North,
        }
    }

    fn offset(&self) -> (i32, i32) {
        match self {
            Direction::North => (-1, 0),
            Direction::East => (0, 1),
            Direction::South => (1, 0),
            Direction::West => (0, -1),
        }
    }
}

fn parse_input(input: &str) -> (Vec<Vec<char>>, (usize, usize), Direction) {
    let mut grid = Vec::new();
    let mut start = (0, 0);
    let mut facing = Direction::North;

    for (row_idx, line) in input.lines().enumerate() {
        let mut row = Vec::new();
        for (col_idx, ch) in line.chars().enumerate() {
            if ch == '^' {
                start = (row_idx, col_idx);
                facing = Direction::North;
                row.push('.');
            } else {
                row.push(ch);
            }
        }
        grid.push(row);
    }

    (grid, start, facing)
}

fn in_bounds(grid: &[Vec<char>], r: i32, c: i32) -> bool {
    r >= 0 && r < grid.len() as i32 && c >= 0 && c < grid[0].len() as i32
}

fn print_grid(grid: &[Vec<char>], position: (i32, i32), visited: &HashSet<(usize, usize)>) {
    for (r, row) in grid.iter().enumerate() {
        for (c, &cell) in row.iter().enumerate() {
            if (r as i32, c as i32) == position {
                print!("G");
            } else if visited.contains(&(r, c)) {
                print!("X");
            } else {
                print!("{}", cell);
            }
        }
        println!();
    }
    println!();
}

fn simulate_patrol(grid: &[Vec<char>], start: (usize, usize), facing: Direction) -> HashSet<(usize, usize)> {
    let mut visited = HashSet::new();
    let mut position = (start.0 as i32, start.1 as i32);
    let mut direction = facing;

    while in_bounds(grid, position.0, position.1) {
        visited.insert((position.0 as usize, position.1 as usize));
        // print_grid(grid, position, &visited);

        let (dr, dc) = direction.offset();
        let next = (position.0 + dr, position.1 + dc);

        if !in_bounds(grid, next.0, next.1) {
            break;
        }

        if in_bounds(grid, next.0, next.1) && grid[next.0 as usize][next.1 as usize] != '#' {
            position = next;
        } else {
            direction = direction.turn_right();
        }
    }

    visited
}

fn test_obstruction(grid: &mut Vec<Vec<char>>, start: (usize, usize), facing: Direction, r: usize, c: usize) -> bool {
    if grid[r][c] == '#' || (r, c) == start {
        return false;
    }

    grid[r][c] = '#';

    let mut visited = HashMap::new();
    let mut position = (start.0 as i32, start.1 as i32);
    let mut direction = facing;

    while in_bounds(grid, position.0, position.1) {
        let key = (position.0 as usize, position.1 as usize, direction);
        if visited.contains_key(&key) {
            grid[r][c] = '.'; // Restore
            return true;
        }
        visited.insert(key, true);

        let (dr, dc) = direction.offset();
        let next = (position.0 + dr, position.1 + dc);

        if !in_bounds(grid, next.0, next.1) {
            break;
        }

        if in_bounds(grid, next.0, next.1) && grid[next.0 as usize][next.1 as usize] != '#' {
            position = next;
        } else {
            direction = direction.turn_right();
        }

        if visited.len() > grid.len() * grid[0].len() {
            println!("Breaking due to excessive steps.");
            break;
        }
    }

    grid[r][c] = '.'; // Restore
    false
}

const DAY: &str = "day06";
pub fn solution() -> io::Result<()> {
    let file_path = format!("src/{}/data.txt", DAY);

    // Use a buffered reader
    let contents = fs::read_to_string(file_path)?;
    let (mut grid, start, facing) = parse_input(&contents);

    // Part 1
    println!("Starting Part 1 simulation:");
    let visited_positions = simulate_patrol(&grid, start, facing);
    println!("Part 1: {} distinct positions visited", visited_positions.len());

    // Part 2
    println!("Starting Part 2 simulation:");
    let mut loop_count = 0;
    for r in 0..grid.len() {
        for c in 0..grid[0].len() {
            if test_obstruction(&mut grid, start, facing, r, c) {
                loop_count += 1;
            }
        }
    }
    println!("Part 2: {} positions cause a loop", loop_count);

    Ok(())
}
