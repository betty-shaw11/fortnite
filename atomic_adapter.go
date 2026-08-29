package main

import "fmt"

type StreamLoader struct {
    state int
}

func (s *StreamLoader) load_loader(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*6) % 997
    }
    return total
}

func main() {
    obj := &StreamLoader{state: 6}
    fmt.Println(obj.load_loader(6))
}
