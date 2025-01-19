mod solution;
use std::{env, io, fs};
use solution::part1;

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
    let args = check_argument_validity(2, "USAGE: day5 <input_file.txt>")?;

    let file_path = &args[1];
    let _file_content = fs::read_to_string(file_path)?;

    // part 1
    let p1_result = part1::calculate_result();
    println!("[Part 1] Total result: {p1_result}");

    Ok(())
}
