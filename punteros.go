package main

import "fmt"

func intercambia(a , b * int) {
    aux := *a
    *a = *b
    *b = aux
}

func main() {
    var a int
    var b int

    fmt.Print("Ingresa un numero: ")
    fmt.Scan(&a)

    fmt.Print("Ingresa otro numero: ")
    fmt.Scan(&b)

    intercambia(&a, &b)

    fmt.Println(a)
    fmt.Println(b)
}
