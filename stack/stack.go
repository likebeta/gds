package stack

import (
    "fmt"
)

type ArrayStack struct {
    arr []interface{}
}

func (s *ArrayStack) Size() int {
    return len(s.arr)
}

func (s *ArrayStack) Empty() bool {
    return len(s.arr) == 0
}

func (s *ArrayStack) Push(v interface{}) *ArrayStack {
    s.arr = append(s.arr, v)
    return s
}

func (s *ArrayStack) Pop() interface{} {
    sz := len(s.arr)
    v := s.arr[sz-1]
    s.arr = s.arr[:sz-1]
    return v
}

func (s *ArrayStack) Top() interface{} {
    sz := len(s.arr)
    return s.arr[sz-1]
}

func (s *ArrayStack) Peek() interface{} {
    return s.Top()
}

func (s *ArrayStack) Slice() []interface{} {
    return s.arr
}

func (s *ArrayStack) String() string {
    return fmt.Sprintf("%v", s.arr)
}
