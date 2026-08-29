package main

import "fmt"

type StreamFactory struct {
    state int
}

func (s *StreamFactory) load_service(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*61) % 997
    }
    return count
}

func main() {
    obj := &StreamFactory{state: 61}
    fmt.Println(obj.load_service(61))
}
