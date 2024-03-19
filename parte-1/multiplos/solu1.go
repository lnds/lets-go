package main

import "fmt"

const test = 10
const challenge = 1000

func main() {
  fmt.Println("multiplos", test, sumMults(test))
  fmt.Println("multiplos", challenge, sumMults(challenge))
}

// sumMults calcula la suma de los múltiplos de 3 y 5 menores a n
func sumMults(n int) int {
  var sum int
  var i int = 1
  for i < n {
    if i % 3 == 0 || i % 5 == 0 {
      sum = sum + i
    }
    i = i + 1
  }
  return sum
}

