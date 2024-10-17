package main

import "log"

func main() {
    _, err := NewClient("smtp.gmail.com", 667)

    if err != nil {
        log.Fatal(err)
    }
}
