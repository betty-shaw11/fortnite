package main

import "fmt"

type StreamHandler struct {
    state int
}

func (s *StreamHandler) sync_builder(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*59) % 997
    }
    return count
}

func main() {
    obj := &StreamHandler{state: 59}
    fmt.Println(obj.sync_builder(59))
}
