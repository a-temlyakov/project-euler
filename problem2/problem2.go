package main

import "fmt"

func fibonacci() func() int {
    x := 0
    y := 1
    return func() int {
        z := x + y
        x = y
        y = z
        return z
    }
}

func main() {
    f := fibonacci()
    fib_num := f()
    even_sum := 0

    for fib_num < 4000000 {
       fib_num = f()
       if fib_num % 2 == 0 {
            even_sum += fib_num
       }
    }

    fmt.Println(even_sum)
}
