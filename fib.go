// 2019 (c) Imre Szakal
// https://github.com/imreszakal
//
// fib.go
// Calculates the biggest Fibonacci-number represented in float64
//
// Output:
// Time: 230.97µs
// 1. fib: 1
// 2. fib: 1
// 3. fib: 2
// 1476. fib:
// 1306989223763398737545115937039993048536618159419209827158963712804246914958
// 6656713050982721611762517795273838124075551803079743968344369778569623080247
// 3309617042775347304891963181519627287463521203531259388682404883801028462229
// 3993455678848254649341365631154415844303003337887773454383151162230325185546
// 81344

package main

import (
    "fmt"
    "time"
    "strconv"
)

func fib(numbers []float64) {
    var x, y float64 = 1, 1
    for c := 0; c < 1476; c++ {
        numbers[c] = x
        x, y = y, x+y
    }
}


func printer(numbers []float64) {
    for i, v := range numbers {
        if i < 7 || i > 1474 {
            fmt.Printf("%d. fib: %.0f\n", i+1, v)

            vs := strconv.FormatFloat(v, 'f', 0, 64)
            fmt.Printf("Number of digits: %v\n\n", len(vs))
        }
    }
}

func main() {
    start := time.Now()
    numbers := make([]float64, 1476, 1476)
    fib(numbers)
    elapsed := time.Since(start)

    fmt.Printf("\nTime: %v\n\n", elapsed)
    printer(numbers)
}
