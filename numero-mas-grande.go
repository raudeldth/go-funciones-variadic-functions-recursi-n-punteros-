package main

import "fmt"

func grande(args... int) int {
    mayor := args[0]

    for _, v := range args {
        if v > mayor {
            mayor = v
        }
    }
    return mayor
}

func main() {
    a := []int{-11, -4, -2}
    fmt.Println(grande(a...))
	fmt.Println(grande(4, 5, 2))
}
