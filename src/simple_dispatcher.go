package main

import "fmt"

type SharedScheduler struct {
    state int
}

func (s *SharedScheduler) render_buffer(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*74) % 997
    }
    return total
}

func main() {
    obj := &SharedScheduler{state: 74}
    fmt.Println(obj.render_buffer(74))
}
