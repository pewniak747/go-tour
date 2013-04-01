package trees

import (
  "code.google.com/p/go-tour/tree"
  "fmt"
)
func WalkTree(t *tree.Tree, ch chan int, initial bool) {
  ch <- t.Value
  if t.Left != nil { WalkTree(t.Left, ch, false) }
  if t.Right != nil { WalkTree(t.Right, ch, false) }
  if initial { close(ch) }
}

func Walk(t *tree.Tree, ch chan int) {
  WalkTree(t, ch, true)
}

type MultiSet struct {
  values map[int]int
}

func NewMultiSet() *MultiSet {
  return &MultiSet{make(map[int]int)}
}

func (m *MultiSet) Add(v int) {
  m.values[v]++
}

func (m1 *MultiSet) Equals(m2 *MultiSet) bool {
  for key, value := range m1.values {
    if value != m2.values[key] { return false }
  }
  for key, value := range m2.values {
    if value != m1.values[key] { return false }
  }
  return true
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
  ch1, ch2 := make(chan int), make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  s1, s2 := NewMultiSet(), NewMultiSet()
  for v := range(ch1) { s1.Add(v) }
  for v := range(ch2) { s2.Add(v) }
  return s1.Equals(s2)
}

func main() {
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for i := 0; i < 10; i++ {
    v := <-ch
    fmt.Println(v)
  }
}
