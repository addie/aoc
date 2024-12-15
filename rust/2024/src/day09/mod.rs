use std::{fs, io};

const DAY: &str = "day09";
pub fn solution() -> io::Result<()> {
    // let file_path = format!("src/{}/data_sample.txt", DAY);
    let file_path = format!("src/{}/data.txt", DAY);
    let contents = fs::read_to_string(file_path)?;

    part1(contents.clone());
    part2(contents.clone());

    Ok(())
}

fn part1(input: String) {
    let mut mapped_memory = map_memory(input.to_string());
    // println!("Mapped memory: {:?}", mapped_memory);

    compress_memory(&mut mapped_memory);
    // println!("Compressed memory: {:?}", mapped_memory);

    let checksum = calc_checksum(&mut mapped_memory);
    println!("Part 1: {:?}", checksum);
}

fn part2(input: String) {
    let mut mapped_memory = map_memory(input.to_string());
    // println!("Mapped memory: {:?}", mapped_memory);

    compress_memory_2(&mut mapped_memory);
    // println!("Compressed memory: {:?}", mapped_memory);

    let checksum = calc_checksum(&mut mapped_memory);
    println!("Part 2: {:?}", checksum);
}

fn calc_checksum(mapped_memory: &mut Vec<String>) -> i64 {
    let mut checksum: i64 = 0;
    let mut i = 0;
    while i < mapped_memory.len() {
        if mapped_memory[i] != '.'.to_string() {
            checksum += mapped_memory[i].parse::<i64>().unwrap() * i as i64;
        }
        i += 1;
    }
    checksum
}

fn compress_memory(mapped_memory: &mut Vec<String>) {
    compress(mapped_memory);
    compress(mapped_memory);
}

fn compress(mapped_memory: &mut Vec<String>) {
    let mut i = 0;
    let mut j = mapped_memory.len() - 1;
    while i <= j {
        while mapped_memory[i] != "." {
            i += 1;
        }
        while mapped_memory[j] == "." {
            j -= 1;
        }
        mapped_memory.swap(i, j);
        i += 1;
        j -= 1;
    }
}

fn compress_memory_2(mapped_memory: &mut Vec<String>) {
    let mut j = mapped_memory.len() - 1;
    while j > 0 {
        compress_next_block(&mut j, mapped_memory);
    }
}

fn compress_next_block(j: &mut usize, mapped_memory: &mut Vec<String>) {
    while *j > 0 && mapped_memory[*j] == "." {
        *j -= 1;
    }
    let mut id_start = *j;
    let mut id_size = 0;
    while *j > 0 && mapped_memory[*j] == mapped_memory[id_start] {
        id_size += 1;
        *j -= 1;
    }
    id_start -= id_size - 1;
    let mut i = 0;
    while i < *j {
        while i < *j && mapped_memory[i] != "." {
            i += 1;
        }
        let free_block_start = i;
        let mut free_block_size = 0;
        while i < *j && mapped_memory[i] == "." {
            free_block_size += 1;
            i += 1;
            if free_block_size >= id_size {
                for k in 0..id_size {
                    mapped_memory.swap(free_block_start + k, id_start + k);
                }
                return;
            }
        }
    }
}

fn find_free_spans(vec: &Vec<String>) -> Vec<(usize, usize)> {
    let mut free_spans = Vec::new();
    let mut start = None;

    for (i, element) in vec.iter().enumerate() {
        if element == "." {
            if start.is_none() {
                start = Some(i);
            }
        } else if let Some(s) = start {
            free_spans.push((s, i - s)); // Store (start_index, length)
            start = None;
        }
    }

    if let Some(s) = start {
        free_spans.push((s, vec.len() - s)); // Handle trailing free space
    }

    free_spans
}
fn find_id_blocks(vec: &Vec<String>) -> Vec<(usize, usize)> {
    let mut id_blocks = Vec::new();
    let mut start = None;

    let mut curr_element = &vec.last().unwrap().to_string();
    for (i, element) in vec.iter().enumerate().rev() {
        if curr_element != "." && element == curr_element || curr_element == "." && element != "." {
            if start.is_none() {
                start = Some(i);
            }
        } else if let Some(s) = start {
            id_blocks.push((s, s - i)); // Store (start_index, length)
            if element == "." {
                start = None;
            } else {
                start = Some(i);
            }
            curr_element = element;
        }
    }

    if let Some(s) = start {
        id_blocks.push((s, s + 1)); // Handle trailing free space
    }

    id_blocks
}

fn process_blocks(vec: &mut Vec<String>, id_blocks: &mut Vec<(usize, usize)>) {
    let mut free_spans = find_free_spans(vec);
    'outer: for id_block in id_blocks {
        let block_size = id_block.1;
        for (idx, free_span) in free_spans.iter().enumerate() {
            if free_span.1 > id_block.0 {
                break 'outer;
            }
            let span_size = free_span.1;
            if block_size <= span_size {
                let start_span = free_span.0;
                let mut counter = 0;
                while counter < block_size {
                    vec.swap(start_span + counter, id_block.0 - counter);
                    counter += 1;
                }
                free_spans = find_free_spans(vec);
                continue 'outer;
            }
        }
    }
}

fn map_memory(input: String) -> Vec<String> {
    let mut id = 0;
    let mut v: Vec<String> = Vec::new();
    for (i, char) in input.chars().enumerate() {
        let current_number = char.to_digit(10).unwrap();
        if i % 2 == 0 {
            // handle ID
            for i in 0..current_number {
                v.push(id.to_string());
            }
            id += 1;
        } else {
            // handle spaces
            for j in 0..current_number {
                v.push(".".to_string())
            }
        }
    }
    v
}
