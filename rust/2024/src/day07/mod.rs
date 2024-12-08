use std::{fs, io};
use strum_macros::EnumCount;

#[derive(EnumCount, Debug, Clone, Copy)]
enum Operator {
    Add,
    Multiply,
    Concatenate,
}

impl Operator {
    // Apply the operator to two numbers
    fn apply(&self, a: i64, b: i64) -> i64 {
        match self {
            Operator::Add => a + b,
            Operator::Multiply => a * b,
            Operator::Concatenate => (a.to_string() + &b.to_string()).parse::<i64>().unwrap(),
        }
    }

    // Represent the operator as a string for debugging
    fn to_string(&self) -> &'static str {
        match self {
            Operator::Add => "+",
            Operator::Multiply => "*",
            Operator::Concatenate => "||",
        }
    }
}
fn parse_input(line: &str) -> (i64, Vec<i64>) {
    let mut parts = line.split(':');
    let value = parts.next().unwrap().trim().to_string().parse().unwrap();
    let list = parts
        .next()
        .unwrap_or("")
        .split_whitespace()
        .map(|s| s.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();
    (value, list)
}

fn generate_operator_permutations(len: usize) -> Vec<Vec<Operator>> {
    let mut permutations = Vec::new();
    let total_permutations = 3usize.pow(len as u32);

    for i in 0..total_permutations {
        let mut current = Vec::new();
        let mut num = i;

        for _ in 0..len {
            match num % 3 {
                0 => current.push(Operator::Add),
                1 => current.push(Operator::Multiply),
                2 => current.push(Operator::Concatenate),
                _ => unreachable!(),
            }
            num /= 3;
        }

        permutations.push(current);
    }

    permutations
}
const DAY: &str = "day07";
pub fn solution() -> io::Result<()> {
    let file_path = format!("src/{}/data.txt", DAY);
    // let file_path = format!("src/{}/data_sample.txt", DAY);

    // Use a buffered reader
    let contents = fs::read_to_string(file_path)?;

    let mut total_count = 0;
    for line in contents.lines() {
        let (value, list) = parse_input(&line);
        let permutations = generate_operator_permutations(list.len()-1);
        for permutation in &permutations {
            let mut result = list[0];
            for i in 0..permutation.len() {
                let operator = permutation[i];
                result = operator.apply(result, list[i+1]);
            }
            // debug(&list, &permutation, &result);

            if result == value {
                println!("Found a hit {}", value);
                total_count += value;
                break;
            }
        }

    }

    println!("Total count {}", total_count);

    Ok(())
}

fn debug(list: &Vec<i64>, operators: &Vec<Operator>, total: &i64) {
    let symbols: Vec<_> = operators.iter().map(|op| op.to_string()).collect();
    let mut result = Vec::new();
    let max_len = list.len().max(symbols.len());
    for i in 0..max_len {
        if let Some(item) = list.get(i) {
            result.push(item.to_string());
        }
        if i == max_len - 1 {
            result.push("= ".to_owned() + &total.to_string());
            break;
        }
        if let Some(item) = symbols.get(i) {
            result.push(item.to_string());
        }
    }
    println!("{}", result.join(" "));
}
