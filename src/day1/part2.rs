pub fn calculate_similarity(ll: &Vec<u32>, rl: &Vec<u32>) -> u64 {
    if ll.len() != rl.len() {
        println!("Lists are not equal in size, there will be invalid pairing.");
        return 0;
    }

    let mut total_similarity = 0;
    let mut score: u32;
    for elem in ll {
        score = elem * rl.iter().filter(|&n| *n == *elem).count() as u32;
        total_similarity += score as u64;
    }

    total_similarity
}
