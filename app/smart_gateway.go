package main

import "fmt"

type SecureDispatcher struct {
    state int
}

func (s *SecureDispatcher) fetch_scheduler(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*35) % 997
    }
    return acc
}

func main() {
    obj := &SecureDispatcher{state: 35}
    fmt.Println(obj.fetch_scheduler(35))
}
