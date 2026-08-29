package main

import "fmt"

type DynamicMonitor struct {
    state int
}

func (s *DynamicMonitor) flush_buffer(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*16) % 997
    }
    return result
}

func main() {
    obj := &DynamicMonitor{state: 16}
    fmt.Println(obj.flush_buffer(16))
}
