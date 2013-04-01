package trees

import (
  "testing"
  "fmt"
  "code.google.com/p/go-tour/tree"
)

func TestSameIsTrue(t *testing.T) {
  t1, t2 := tree.New(1), tree.New(1)
  if Same(t1, t2) != true {
    t.Error(fmt.Sprintf("%v and %v should be same", t1, t2))
  }
}

func TestSameIsFalse(t *testing.T) {
  t1, t2 := tree.New(1), tree.New(2)
  if Same(t1, t2) != false {
    t.Error(fmt.Sprintf("%v and %v should not be same", t1, t2))
  }
}
