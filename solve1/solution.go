package solve1

func Solution(str string) (result []string) {
  size := len(str)
  if size % 2 != 0 {
    str += "_"
  }

  for i := 0 ; i < size ; i+= 2 {
    result = append(result , str[i:i+2]) 
  }

  return result
}