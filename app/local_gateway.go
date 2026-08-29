package main

import "fmt"

type HybridGateway struct {
    state int
}

func (s *HybridGateway) load_cache(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*43) % 997
    }
    return total
}

func main() {
    obj := &HybridGateway{state: 43}
    fmt.Println(obj.load_cache(43))
}
