// ioライブラリはstd標準ライブラリに入っている
// 標準ライブラリで定義されているアイテムの中のいくつかを、すべてのプログラムのスコープに取り込む
// このセットはprelude(プレリュード：前奏曲)と呼ばれる
use std::{cmp::Ordering, io};

use rand::Rng;

mod training;

#[allow(unreachable_code)]
fn main() {
    training::training();

    return;

    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..101);

    loop {
        println!("Please input your guess.");
    
        let mut guess = String::new();
    
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");
    
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num, 
            Err(_) => continue,
        };

        println!("You guessd: {}", guess);
    
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!!"),
            Ordering::Greater => println!("Too big!!"),
            Ordering::Equal => { 
                println!("You win!!");
                break;
            }
        }
    }
}

