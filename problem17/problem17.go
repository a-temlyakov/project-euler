package main

import "fmt"
import "math"

func main() {
    m := make(map[int]string)
    m[1] = "one"
    m[2] = "two"
    m[3] = "three"
    m[4] = "four"
    m[5] = "five"
    m[6] = "six"
    m[7] = "seven"
    m[8] = "eight"
    m[9] = "nine"
    m[10] = "ten"
    m[11] = "eleven"
    m[12] = "twelve"
    m[13] = "thirteen"
    m[14] = "fourteen"
    m[15] = "fifteen"
    m[16] = "sixteen"
    m[17] = "seventeen"
    m[18] = "eighteen"
    m[19] = "nineteen"
    m[20] = "twenty"
    m[30] = "thirty"
    m[40] = "forty"
    m[50] = "fifty"
    m[60] = "sixty"
    m[70] = "seventy"
    m[80] = "eighty"
    m[90] = "ninety"

    sum := 0;
    for i := 1; i <= 20; i++ {
        sum += len(m[i])
    }

    for i := 21; i < 100; i++ { 
        a := i - int(math.Floor(float64(i) / 10.0)) * 10

        if a == 0 {
            sum += len(m[i])
        } else {
            sum += len(m[i - a]) + len(m[a])
            m[i] = m[i - a] + m[a]
        }
    }

    sum += len("onehundred")
    m[100] = "onehundred"

    for i := 101; i < 1000; i++ {
        const hundred = "hundred"
        b := int(math.Floor(float64(i) / 100.0))
        c := i - b * 100

        if c == 0 {
            m[i] = m[b] + hundred
            sum += len(m[b]) + len(hundred)
        } else {
            m[i] = m[b] + hundred + "and" + m[c]
            sum += len(m[b]) + len(hundred) + len("and") + len(m[c])
        }
    }

    sum += len("onethousand")
    m[1000] = "onethousand"

    fmt.Println(sum)
}
