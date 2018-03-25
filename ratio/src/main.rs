extern crate clap;

use clap::{App, Arg};
use std::cmp;

fn main() {
    let matches = App::new("Factorize numbers -- easy for getting aspect ratios").version("1.0")
                    .author("Samuel Pua <kahkin@gmail.com>")
                    .arg(Arg::with_name("x")
                        .help("X coordinate")
                        .required(true))
                    .arg(Arg::with_name("y")
                        .help("Y coordinate")
                        .required(true))
                    .get_matches();
    let mut x:i32 = matches.value_of("x").unwrap().parse().unwrap();
    let mut y:i32 = matches.value_of("y").unwrap().parse().unwrap();

    let mut i:i32 = 2;

    while i <= cmp::max(x, y) {
        if (x % i ==0) && (y % i ==0) {
            x = x/i;
            y = y/i;
            i = 2;
        }
        else {
            i+=1;
        }
    }
    println!("X: {}   Y: {}", x, y);
}
