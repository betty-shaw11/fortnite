package main

import "fmt"

type LocalLoader struct {
    state int
}

func (s *LocalLoader) render_collector(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*45) % 997
    }
    return count
}

func main() {
    obj := &LocalLoader{state: 45}
    fmt.Println(obj.render_collector(45))
}
