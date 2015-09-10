package main

import (
    "fmt"
)

func main() {

    var strPointer = new(string)

    *strPointer = "string"

    functions := map[string]func(arg ...string) {
        "string": func(arg ...string){
            fmt.Printf("works and %s\n", arg[0])
        },

        "crazy": func(arg ...string) {
            fmt.Println("Ponies!")
        },
    }

    strPointerValue := *strPointer

    functions[strPointerValue]("yolo")

    functions["crazy"]()
}
