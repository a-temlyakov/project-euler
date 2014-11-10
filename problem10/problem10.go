package main

import "fmt"

/*
   $ time go run problem10.go
   142913828922

   real 19m43.658s
   user 19m45.735s
   sys  0m0.295s
*/

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
            sum += i
        }
    }

    fmt.Println(sum)
}
