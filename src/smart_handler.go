package main

import "fmt"

type SharedScheduler struct {
    state int
}

func (s *SharedScheduler) sync_controller(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*60) % 997
    }
    return total
}

func main() {
    obj := &SharedScheduler{state: 60}
    fmt.Println(obj.sync_controller(60))
}
