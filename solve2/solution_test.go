package solve2

// import testing package
import (
  "testing"
)

func TestSolution1(t *testing.T) {
  got := []int{2,3,1}
  want := 0
  result := Solution(got)
  if int(result) != want {
    // report mismatch
    t.Errorf("\n Solution() = \n got %v, \n want %v , \n result = %v", got, want, result)
  }
}

func TestSolution2(t *testing.T) {
  got := []int{5,10,14}
  want := 1
  result := Solution(got)
  if result != want {
    // report mismatch
    t.Errorf("Solution() = %q, want %q", got, want)
  }
}

