package main

import (
  "fmt"
  "main/solve1"
  "main/solve2"
)

func main (){
  fmt.Println("Hello")
  fmt.Println(solve1.Solution("abc"))

  fmt.Println(solve2.Solution([]int{2,3,1}))
  fmt.Println(solve2.Solution([]int{5,10,14}))
}
