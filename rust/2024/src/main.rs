#![allow(dead_code)]
#![allow(unused_variables)]

use std::env;
use std::process::exit;

macro_rules! mod_list {
    ($($name:ident),*) => {
        $(
            mod $name;
        )*
    };
}

mod_list!(
    day01, day02, day03, day04, day05, day06, day07, day08, day09, day10, day11, day12, day13,
    day14, day15, day16, day17, day18, day19, day20, day21, day22, day23, day24, day25
);

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        println!("No args specified");
        exit(1);
    }
    match args[1].as_str() {
        "1" => day01::solution().unwrap(),
        "2" => day02::solution().unwrap(),
        "3" => day03::solution().unwrap(),
        "4" => day04::solution().unwrap(),
        "5" => day05::solution().unwrap(),
        "6" => day06::solution().unwrap(),
        "7" => day07::solution().unwrap(),
        "8" => day08::solution().unwrap(),
        "9" => day09::solution().unwrap(),
        "10" => day10::solution().unwrap(),
        "11" => day11::solution().unwrap(),
        "12" => day12::solution().unwrap(),
        "13" => day13::solution().unwrap(),
        "14" => day14::solution().unwrap(),
        "15" => day15::solution().unwrap(),
        "16" => day16::solution().unwrap(),
        "17" => day17::solution().unwrap(),
        "18" => day18::solution().unwrap(),
        "19" => day19::solution().unwrap(),
        "20" => day20::solution().unwrap(),
        "21" => day21::solution().unwrap(),
        "22" => day22::solution().unwrap(),
        "23" => day23::solution().unwrap(),
        "24" => day24::solution().unwrap(),
        "25" => day25::solution().unwrap(),
        _ => {}
    }
}
