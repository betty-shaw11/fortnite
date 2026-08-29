package main

import "fmt"

type SharedRegistry struct {
    state int
}

func (s *SharedRegistry) render_scheduler(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*65) % 997
    }
    return value
}

func main() {
    obj := &SharedRegistry{state: 65}
    fmt.Println(obj.render_scheduler(65))
}
