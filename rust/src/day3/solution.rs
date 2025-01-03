pub mod part1 {
use regex::Regex;

pub fn calculate_result(input: String) -> usize {
    let re = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").unwrap();
    let pairs_found: Vec<(usize, usize)> = re.captures_iter(&input).map(
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
