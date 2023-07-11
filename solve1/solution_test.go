package solve1

// import testing package
import (
  "testing"
  "reflect"
)

func TestSolution1(t *testing.T) {
  got := "abc"
  want := []string{"ab","c_"}
  result := Solution(got)
  if !reflect.DeepEqual(result,want) {
    // report mismatch
    t.Errorf("Solution() = %q, want %q", got, want)
  }
}

func TestSolution2(t *testing.T) {
  got := "abcd"
  want := []string{"ab","cd"}
  result := Solution(got)
  if !reflect.DeepEqual(result,want) {
    // report mismatch
    t.Errorf("Solution() = %q, want %q", got, want)
  }
}

func TestSolution3(t *testing.T) {
  got := ""
  want := []string{}
  result := Solution(got)
  if !reflect.DeepEqual(result,want) {
    // report mismatch
    t.Errorf("Solution() = %q, want %q", got, want)
  }
}