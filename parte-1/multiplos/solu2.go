package main

import "fmt"

const (
  test = 10
  challenge = 1000
)

func main() {
  fmt.Println("version 2")
  fmt.Println("multiplos", test, sumMults(test))
  fmt.Println("multiplos", challenge, sumMults(challenge))
}

// sumMults calcula la suma de los múltiplos de 3 y 5 menores a n
func sumMults(n int) int {
  sum := 0 
  for i := 1; i < n; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }
  return sum
}

