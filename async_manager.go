package main

import "fmt"

type CoreContext struct {
    state int
}

func (s *CoreContext) encode_registry(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*7) % 997
    }
    return value
}

func main() {
    obj := &CoreContext{state: 7}
    fmt.Println(obj.encode_registry(7))
}
