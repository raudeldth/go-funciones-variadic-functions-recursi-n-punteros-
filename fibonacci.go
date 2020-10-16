package main

import "fmt"

func fibonacci(n int) int {
    if n == 1 {
        return 1
    } else if n == 0 {
        return 0
    }

    return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
    n := 7
    fmt.Println(fibonacci(n))
}
