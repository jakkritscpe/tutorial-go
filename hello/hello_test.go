// file hello_test.go
package hello

// import testing package
import "testing"

func TestHello(t *testing.T) {
  want := "Hello World"
  if got := Hello(); got != want {
    // report mismatch
    t.Errorf("Hello() = %q, want %q", got, want)
  }
}