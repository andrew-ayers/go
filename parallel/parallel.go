package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    runtime.GOMAXPROCS(4)

    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Starting Go Routines")
    go func() {
        defer wg.Done()

        for char := 'a'; char < 'a'+26; char++ {
            fmt.Printf("%c ", char)
            //time.Sleep(time.Duration(random(1, 5)) * time.Second)
        }
    }()

    go func() {
        defer wg.Done()

        for number := 1; number < 27; number++ {
            fmt.Printf("%d ", number)
            //time.Sleep(time.Duration(random(1, 5)) * time.Second)
        }
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()

    fmt.Println("\nTerminating Program")

    //time.Sleep(60 * time.Second)
}