pub fn training() {
   let s = String::from("hello world");
   let word = first_word(&s);

   let first: &str = &s[0..word];

   println!("first word is '{}'", first);
}

fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}
