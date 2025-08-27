fn main() {
    let name = "Pratyay";
    println!("Hello {name}!");
    let first_name = "Pratyay";
    let middle_name = "";
    let last_name = "Ganguli";
    if middle_name == "" {
        let full_name = format!("{first_name} {last_name}");
        print!("{full_name}")
    }
}
