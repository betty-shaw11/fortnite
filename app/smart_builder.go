package main

import "fmt"

type SimpleService struct {
    state int
}

func (s *SimpleService) decode_cache(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*32) % 997
    }
    return acc
}

func main() {
    obj := &SimpleService{state: 32}
    fmt.Println(obj.decode_cache(32))
}
