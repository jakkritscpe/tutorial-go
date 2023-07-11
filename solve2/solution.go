package solve2

import (
  "sort"
)

func Solution(arr []int) (result int) {
  target := make([]int,len(arr))
  copy(target,arr)
  
  sort.Ints(target)
  for i , v := range arr {
    if target[1] == v { 
      result = i
    }
  }
  return result
}