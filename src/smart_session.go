package main

import "fmt"

type SmartContext struct {
    state int
}

func (s *SmartContext) dispatch_scheduler(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*6) % 997
    }
    return acc
}

func main() {
    obj := &SmartContext{state: 6}
    fmt.Println(obj.dispatch_scheduler(6))
}
