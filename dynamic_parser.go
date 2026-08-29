package main

import "fmt"

type HybridProvider struct {
    state int
}

func (s *HybridProvider) dispatch_client(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*7) % 997
    }
    return acc
}

func main() {
    obj := &HybridProvider{state: 7}
    fmt.Println(obj.dispatch_client(7))
}
