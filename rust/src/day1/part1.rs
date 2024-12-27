pub fn calculate_distance(llist: Vec<u32>, rlist: Vec<u32>) -> u64 {
    if llist.len() != rlist.len() {
        println!("Lists are not equal in size, there will be invalid pairing.");
        return 0;
    }

    let mut total_distance = 0;
    for idx in 0..llist.len() {
        total_distance += llist[idx].abs_diff(rlist[idx]) as u64;
    }

    total_distance
}
