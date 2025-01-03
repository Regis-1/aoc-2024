use regex::Regex;
pub mod part1 {
use super::Regex;

pub fn calculate_result(input: &str) -> usize {
    let re = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").unwrap();
    let pairs_found: Vec<(usize, usize)> = re.captures_iter(input).map(
        |pair| {
            let n1: usize = pair[1].parse().unwrap();
            let n2: usize = pair[2].parse().unwrap();
            (n1, n2)
        }
    ).collect();

    let mut total: usize = 0;
    for &pair in pairs_found.iter() {
        total += pair.0 * pair.1;
    }

    total
}
}

pub mod part2 {
use super::Regex;

pub fn calculate_result(input: &str) -> usize {
    let pattern = r"mul\((\d{1,3}),(\d{1,3})\)|(don't\(\))|(do\(\))";
    let re = Regex::new(pattern).unwrap();
    let mut do_include: bool = true;

    let mut total: usize = 0;
    for capture in re.captures_iter(input) {
        if let Some(num1) = capture.get(1) {
            if do_include {
                let num1: usize = num1.as_str().parse().unwrap();
                let num2: usize = capture.get(2).unwrap().as_str().parse().unwrap();
                total += num1 * num2;
            }
        }
        else {
            let instruction = capture.get(0).unwrap().as_str();
            if instruction == "do()" { do_include = true; }
            else if instruction == "don't()" { do_include = false; }
        }
    }

    total
}
}
