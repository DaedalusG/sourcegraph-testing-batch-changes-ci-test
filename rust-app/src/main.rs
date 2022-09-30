fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_good() {
        assert_eq!(3, 3);
    }

    #[test]
    fn test_bad() {
        assert_eq!(1, 3);
    }
}
