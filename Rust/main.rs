extern crate reqwest;
extern crate regex;

use regex::Regex;

fn get_data(s: &str) -> String {
    
    let mut info: String = String::new();
    
    match reqwest::get(s) {
        Ok(mut response) => {
            // Check if 200 OK
            if response.status() == reqwest::StatusCode::OK {
                match response.text() {
                    Ok(text) => info = text,
                    Err(..) => println!("Could not read the response text")
                }
            } else {
                println!("Response was not 200 OK");
            }
        }
        Err(..) => println!("Could not make the request")
    }
    info
}


fn main() {
    
    let mut orders: Vec<String> = Vec::new(); // to store all orders in backlog

    // download the data
    let data = get_data("https://orders.olkb.com");
    // compile regexp - with 4 extra chars to avoid counting merged orders
    let re = Regex::new(r"<li>10000\d{4}").unwrap();
    // process the data
    for cap in re.captures_iter(&data) {
        let tobesliced = &cap[0];       // get data from regexp
        let slice = &tobesliced[4..];   // eliminate 4 extrachars before loading into vector
        orders.push(slice.to_string());
    }

    let order_number = "100007000"; // put your own order number
    let order_position = orders.iter().position(|r| r == order_number).unwrap() + 1;
    println!("olkb position: {}", order_position);
    println!("------------------");
    println!("Total orders: {}", orders.len());
    println!("Order #: {}", order_number);
    
}