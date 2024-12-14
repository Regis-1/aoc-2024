use std::{env, fs, io};
mod part1;
mod part2;

fn get_vectors_from_input(content: String) -> (Vec<u32>, Vec<u32>) {
    let mut ll: Vec<u32> = vec![];
    let mut rl: Vec<u32> = vec![];

    for line in content.lines() {
        let number_pair: Vec<&str> = line.split_whitespace().collect();
        ll.push(number_pair[0].parse().unwrap());
        rl.push(number_pair[1].parse().unwrap());
    }

    (ll, rl)
}

fn check_argument_validity(arg_num: usize, usage_info: &str) -> Result<Vec<String>, io::Error> {
    let args: Vec<String> = env::args().collect();
    if args.len() != arg_num {
        let err = Err(
            io::Error::new(io::ErrorKind::InvalidInput, "Invalid number of arguments!")
        );
        println!("{}", usage_info);
        return err;
    }

    Ok(args)
}

fn main() -> Result<(), io::Error> {
    let args = check_argument_validity(2, "USAGE: day1 <input_file.txt>")?;

    let file_path = &args[1];
    let file_content = fs::read_to_string(file_path)?;
    let (left_list, right_list) = get_vectors_from_input(file_content);

    // part 1
    let mut part1_llist = left_list.clone();
    part1_llist.sort();
    let mut part1_rlist = right_list.clone();
    part1_rlist.sort();

    let distance = part1::calculate_distance(part1_llist, part1_rlist);
    println!("[Part 1] Total distance: {}", distance);

    // part 2
    let similarity = part2::calculate_similarity(&left_list, &right_list);
    println!("[Part 2] Total similarity: {}", similarity);

    Ok(())
}
