package main

import "fmt"

func isPrime(x int) bool {
    if x == 0 || x == 1 {
        return false
    }

    for i := 2; i < x; i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}

func main() {
    sum := 0
    for i := 0; i < 2000000; i++ {
        if isPrime(i) {
            fmt.Println(i)
            sum += i
        }
    }

    fmt.Println(sum)
}
