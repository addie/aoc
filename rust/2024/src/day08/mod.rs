use std::collections::HashMap;
use std::{fs, io};

fn parse_input(reader: &str) -> Vec<Vec<char>> {
    let mut grid: Vec<Vec<char>> = Vec::new();
    for line in reader.lines() {
        let line = line.trim();
        let row = line.chars().collect::<Vec<char>>();
        grid.push(row);
    }
    grid
}

const DAY: &str = "day08";
pub fn solution() -> io::Result<()> {
    // let file_path = format!("src/{}/data_sample.txt", DAY);
    let file_path = format!("src/{}/data.txt", DAY);
    let contents = fs::read_to_string(file_path)?;
    let grid = parse_input(&*contents);
    print_grid(&grid);

    part1(grid.clone());
    part2(grid.clone());

    Ok(())
}

fn part1(mut grid: Vec<Vec<char>>) {
    let mut m: HashMap<char, Vec<(usize, usize)>> = HashMap::new();
    for (r, row) in grid.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            if cell == &'.' {
                continue;
            }
            m.entry(cell.clone()).or_insert_with(Vec::new).push((r, c));
        }
    }
    for key in m.keys() {
        let frequencies = m.get(key).unwrap();
        for i in 0..frequencies.len() {
            for j in i + 1..frequencies.len() {
                let (f1r, f1c, f2r, f2c) = (
                    frequencies[i].0 as i32,
                    frequencies[i].1 as i32,
                    frequencies[j].0 as i32,
                    frequencies[j].1 as i32,
                );
                println!("freq pair ({},{}), ({},{})", f1r, f1c, f2r, f2c);
                let a1 = get_antinode((f1r, f1c), get_distance((f1r, f1c), (f2r, f2c)));
                let a2 = get_antinode((f2r, f2c), get_distance((f2r, f2c), (f1r, f1c)));
                if a1.0 >= 0 && a1.0 < grid.len() as i32 && a1.1 >= 0 && a1.1 < grid[0].len() as i32 {
                    grid[a1.0 as usize][a1.1 as usize] = '#';
                }
                if a2.0 >= 0 && a2.0 < grid.len() as i32 && a2.1 >= 0 && a2.1 < grid[0].len() as i32 {
                    grid[a2.0 as usize][a2.1 as usize] = '#';
                }
            }
        }
    }
    print_grid(&grid);
    let mut count = 0;
    for row in grid {
        for cell in row {
            if cell == '#' {
                count += 1;
            }
        }
    }
    println!("part1: {}", count);
}

fn part2(mut grid: Vec<Vec<char>>) {
    let mut m: HashMap<char, Vec<(usize, usize)>> = HashMap::new();
    for (r, row) in grid.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            if cell == &'.' {
                continue;
            }
            m.entry(cell.clone()).or_insert_with(Vec::new).push((r, c));
        }
    }
    for key in m.keys() {
        let frequencies = m.get(key).unwrap();
        for i in 0..frequencies.len() {
            for j in i + 1..frequencies.len() {
                let (f1r, f1c, f2r, f2c) = (
                    frequencies[i].0 as i32,
                    frequencies[i].1 as i32,
                    frequencies[j].0 as i32,
                    frequencies[j].1 as i32,
                );
                println!("freq pair ({},{}), ({},{})", f1r, f1c, f2r, f2c);
                mark_all_antinodes(&mut grid, (f1r, f1c), get_distance((f1r, f1c), (f2r, f2c)));
            }
        }
    }
    print_grid(&grid);
    let mut antinode_count = 0;
    for r in grid {
        for cell in r {
            if cell != '.' {
                antinode_count += 1;
            }
        }
    }
    println!("part2: {}", antinode_count);
}

fn get_distance(f1: (i32, i32), f2: (i32, i32)) -> (i32, i32) {
    (f2.0 - f1.0,f2.1 - f1.1)
}
fn mark_all_antinodes(
    grid: &mut Vec<Vec<char>>,
    start: (i32, i32),
    dist: (i32, i32),
) {
    let rows: i32 = grid.len() as i32;
    let cols: i32 = grid[0].len() as i32;
    let (mut x1, mut y1) = start;
    let (dr, dc) = dist;
    loop {
        (x1, y1) = (x1+dr, y1+dc);
        if !is_valid((x1, y1), rows, cols) {
            break;
        }
        grid[x1 as usize][y1 as usize] = '#';
    }
    let (mut x1, mut y1) = start;
    let (dr, dc) = (dist.0*-1, dist.1*-1);
    loop {
        (x1, y1) = (x1+dr, y1+dc);
        if !is_valid((x1, y1), rows, cols) {
            break;
        }
        grid[x1 as usize][y1 as usize] = '#';
    }
}

fn is_valid(coord: (i32, i32), rows: i32, cols: i32) -> bool {
    let (r, c) = coord;
    r >= 0 && r < rows && c >= 0 && c < cols
}

fn get_antinode(
    start: (i32, i32),
    dist: (i32, i32),
) -> (i32, i32) {
    let (x1, y1) = start;
    let (dr, dc) = dist;
    (x1 + -1*dr, y1 + -1*dc)
}

fn print_grid(grid: &Vec<Vec<char>>) {
    let count = 4;
    print!("    ");
    for i in 0..grid.len() {
        print!("{}", i);
        let spaces = " ".repeat(count - (i as i32).abs().to_string().len());
        print!("{}", spaces);
    }
    println!();
    for (i, row) in grid.iter().enumerate() {
        print!("{}", i);
        let spaces = " ".repeat(count - (i as i32).abs().to_string().len());
        print!("{}", spaces);
        for cell in row {
            print!("{}   ", cell);
        }
        println!();
    }
}
