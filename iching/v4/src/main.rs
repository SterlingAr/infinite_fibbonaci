use rand::Rng;
use std::{thread, time::Duration};

const MAX_THREADS: usize = 13;

fn generate_random_matrix_rain() {
    let yin_yang = ["1", "0"];
    let mut rng = rand::thread_rng();

    loop {
        // Randomly choose between printing a single digit, season, trigram, or hexagram
        let choice = rng.gen_range(0..4);
        let output = match choice {
            0 => yin_yang[rng.gen_range(0..2)].to_string(),          // Single digit
            1 => generate_binary_season(&yin_yang, &mut rng),        // Season
            2 => generate_binary_trigram(&yin_yang, &mut rng),       // Trigram
            _ => generate_binary_hexagram(&yin_yang, &mut rng),      // Hexagram
        };

        print!("{}", output);
        let sleep_time = rng.gen_range(50..200); // Random sleep to create a staggered effect
        thread::sleep(Duration::from_millis(sleep_time));
    }
}

fn generate_binary_season(yin_yang: &[&str; 2], rng: &mut rand::rngs::ThreadRng) -> String {
    let season = (0..2).map(|_| yin_yang[rng.gen_range(0..2)]).collect();
    season
}

fn generate_binary_trigram(yin_yang: &[&str; 2], rng: &mut rand::rngs::ThreadRng) -> String {
    let trigram = (0..3).map(|_| yin_yang[rng.gen_range(0..2)]).collect();
    trigram
}

fn generate_binary_hexagram(yin_yang: &[&str; 2], rng: &mut rand::rngs::ThreadRng) -> String {
    let hexagram = (0..6).map(|_| yin_yang[rng.gen_range(0..2)]).collect();
    hexagram
}

fn main() {
    let mut handles = vec![];

    for _ in 0..MAX_THREADS {
        let handle = thread::spawn(move || {
            generate_random_matrix_rain();
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
}
