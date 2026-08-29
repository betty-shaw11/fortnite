package main

import "fmt"

type StreamProvider struct {
    state int
}

func (s *StreamProvider) flush_scheduler(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*57) % 997
    }
    return count
}

func main() {
    obj := &StreamProvider{state: 57}
    fmt.Println(obj.flush_scheduler(57))
}
