use std::collections::{HashMap, HashSet, VecDeque};
use std::fs;
use std::io;

const DAY: &str = "day05";

pub fn solution() -> io::Result<()> {
    let file_path = format!("src/{}/data.txt", DAY);

    // Use a buffered reader
    let contents = fs::read_to_string(file_path)?;

    let parts: Vec<&str> = contents.split("\n\n").collect();

    let rankings: Vec<(i32, i32)> = parts[0]
        .lines()
        .map(|line| {
            let nums: Vec<i32> = line
                .split('|')
                .map(|s| s.trim().parse().expect("invalid number"))
                .collect();
            (nums[0], nums[1])
        }).collect();
    let entrees: Vec<Vec<i32>> = parts[1]
        .lines()
        .map(|line| {
            line.split(',')
                .map(|s| s.trim().parse().expect("invalid number"))
                .collect()
        }).collect();

    let mut sum_middle_1 = 0;
    let mut rules: HashMap<i32, HashSet<i32>> = HashMap::new();

    // Part 1
    for rank in rankings.as_slice() {
        // Use entry to handle missing keys
        rules.entry(rank.0).or_insert_with(HashSet::new).insert(rank.1);
    }
    for entree in entrees.as_slice() {
        if is_correctly_ordered(&entree, &rules) {
            if !entree.is_empty() {
                let mid = entree[entree.len() / 2];
                sum_middle_1 += mid;
            }
        }
    }

    // Part 2
    let mut rules: HashMap<i32, HashSet<i32>> = HashMap::new();
    for rule in rankings {
        rules.entry(rule.0).or_default().insert(rule.1);
    }

    let mut sum_middle_2 = 0;
    for entree in entrees.as_slice() {
        if is_correctly_ordered(&entree, &rules) {
            continue;
        }

        let fixed_order = fix_order(&entree, &rules);
        if !fixed_order.is_empty() {
            let middle = fixed_order[fixed_order.len() / 2];
            sum_middle_2 += middle;
        }
    }

    println!("Day05 Part1: {}", sum_middle_1);
    println!("Day05 Part2: {}", sum_middle_2);
    Ok(())
}
fn is_correctly_ordered(update: &[i32], rules: &HashMap<i32, HashSet<i32>>) -> bool {
    for i in 0..update.len() {
        for j in i + 1..update.len() {
            if let Some(depends_on) = rules.get(&update[j]) {
                if depends_on.contains(&update[i]) {
                    return false;
                }
            }
        }
    }
    true
}

fn fix_order(update: &[i32], rules: &HashMap<i32, HashSet<i32>>) -> Vec<i32> {
    let mut in_degree = HashMap::new();
    let mut graph: HashMap<i32, Vec<i32>> = HashMap::new();

    // Build the graph and calculate in-degrees
    for &page in update {
        in_degree.entry(page).or_insert(0);
    }
    for (&start, ends) in rules {
        for &end in ends {
            if update.contains(&start) && update.contains(&end) {
                graph.entry(start).or_default().push(end);
                *in_degree.entry(end).or_insert(0) += 1;
            }
        }
    }

    // Topological sort using Kahn's Algorithm
    let mut queue: VecDeque<i32> = in_degree
        .iter()
        .filter(|&(_, &deg)| deg == 0)
        .map(|(&page, _)| page)
        .collect();
    let mut sorted = Vec::new();

    while let Some(page) = queue.pop_front() {
        sorted.push(page);
        if let Some(neighbors) = graph.get(&page) {
            for &neighbor in neighbors {
                if let Some(deg) = in_degree.get_mut(&neighbor) {
                    *deg -= 1;
                    if *deg == 0 {
                        queue.push_back(neighbor);
                    }
                }
            }
        }
    }

    if sorted.len() != update.len() {
        return Vec::new(); // Return empty if there was a cycle (shouldn't happen in valid input)
    }
    sorted
}