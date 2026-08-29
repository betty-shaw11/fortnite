package main

import "fmt"

type StreamHandler struct {
    state int
}

func (s *StreamHandler) run_cache(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*89) % 997
    }
    return value
}

func main() {
    obj := &StreamHandler{state: 89}
    fmt.Println(obj.run_cache(89))
}
