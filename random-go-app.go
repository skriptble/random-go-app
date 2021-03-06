package main

import (
    "fmt"
    "strconv"
    "time"
)

func doStuff(c chan string) {
    mega_awesome := <- c
    fmt.Println("Mega Awesome: " + mega_awesome)
}

func doinThoseProcesses(c chan string) {
    something := <- c
    fmt.Println("Communicating:" + something)
}
func main() {
    c := make(chan string)
    for i := 0; i < 100; i++ {
        go doStuff(c)
    }
    for j := 0; j < 100; j++ {
        c <- "whoa " + strconv.Itoa(j)
        time.Sleep(100 * time.Millisecond)
    }
    fmt.Println("Hello World! I AM AWESOME")
    fmt.Println("More awesome than you know!")
    fmt.Println("I actually love writing Go!")
    fmt.Println("Why you ask? Because it is AWESOME!")
    fmt.Println("Now this is a story all about how my life got flipped-turned upside down.")
    fmt.Println("If everything relies on context, how can we understand anything about anything?")


}
