#![allow(dead_code)]
#![allow(unused_variables)]

macro_rules! mod_list {
    ($($name:ident),*) => {
        $(
            mod $name;
        )*
    };
}

mod_list!(day01, day02, day03, day04, day05, day06, day07);

fn main() {
    // day01::solution().unwrap();
    // day02::solution().unwrap();
    // day03::solution().unwrap();
    // day04::solution().unwrap();
    // day05::solution().unwrap();
    // day06::solution().unwrap();
    day07::solution().unwrap();
}
