use std::collections::HashMap;
use std::{fs, io};

const DAY: &str = "day11";
pub fn solution() -> io::Result<()> {
    // let file_path = format!("src/{}/data_sample.txt", DAY);
    let file_path = format!("src/{}/data.txt", DAY);

    // Use a buffered reader
    let contents = fs::read_to_string(file_path)?;
    let init: Vec<i64> = contents.split(' ').map(|v| v.parse().unwrap()).collect();

    println!("Part 1 {}", part1(init.clone(), 25));
    println!("Part 2 {}", part2(init.clone(), 75));

    Ok(())
}

// - If the stone is engraved with the number 0, it is replaced by a stone
// engraved with the number 1.
// - If the stone is engraved with a number that has an even number of
// digits, it is replaced by two stones. The left half of the digits are
// engraved on the new left stone, and the right half of the digits are
// engraved on the new right stone. (The new numbers don't keep extra
// leading zeroes: 1000 would become stones 10 and 0.)
// - If none of the other rules apply, the stone is replaced by a new
// stone; the old stone's number multiplied by 2024 is engraved on the
// new stone.

fn part1(mut init: Vec<i64>, num_blinks: usize) -> usize {
    for i in 0..num_blinks {
        // println!("number of blinks: {}", i);
        let mut result = Vec::new();
        for num in &init {
            match num {
                0 => result.push(1),
                n if n.abs().to_string().len() % 2 == 0 => {
                    let s = n.abs().to_string();
                    let mid = s.len() / 2;
                    let (left, right) = s.split_at(mid);
                    if let (Ok(left_num), Ok(right_num)) =
                        (left.parse::<i64>(), right.parse::<i64>())
                    {
                        if n < &0 {
                            result.push(-left_num);
                            result.push(-right_num);
                        } else {
                            result.push(left_num);
                            result.push(right_num);
                        }
                    }
                }
                n => result.push(n * 2024),
            }
        }
        // println!("{:?}", result);
        init = result.clone();
    }
    init.len()
}
fn part2(init: Vec<i64>, num_blinks: i32) -> i64 {
    let mut cache: HashMap<i32, HashMap<i64, i64>> = HashMap::new();

    // Init cache
    for r in 1..=num_blinks {
        cache.insert(r, HashMap::new());
    }

    // Add the results of count_stones for each stone in the list
    init
        .into_iter()
        .map(|stone| count_stones(stone, &mut cache, num_blinks))
        .sum()
}

fn count_stones(stone: i64, cache: &mut HashMap<i32, HashMap<i64, i64>>, blink_num: i32) -> i64 {
    if blink_num == 0 {
        return 1;
    }

    // Check if the value is already cached in a limited scope.
    {
        let iter_cache = cache.entry(blink_num).or_insert_with(HashMap::new);
        if let Some(&val) = iter_cache.get(&stone) {
            return val;
        }
    }

    // At this point, the borrow on iter_cache has ended.
    // Compute val without holding the cache borrow.
    let str_stone = stone.to_string();
    let val = if stone == 0 {
        count_stones(1, cache, blink_num - 1)
    } else if str_stone.len() % 2 == 0 {
        let mid = str_stone.len() / 2;
        let first: i64 = str_stone[..mid].parse().unwrap();
        let second: i64 = str_stone[mid..].parse().unwrap();
        count_stones(first, cache, blink_num - 1) + count_stones(second, cache, blink_num - 1)
    } else {
        count_stones(stone * 2024, cache, blink_num - 1)
    };

    // Now borrow iter_cache again to store the computed value.
    let iter_cache = cache.entry(blink_num).or_insert_with(HashMap::new);
    iter_cache.insert(stone, val);
    val
}