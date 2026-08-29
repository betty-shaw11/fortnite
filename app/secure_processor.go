package main

import "fmt"

type AsyncHandler struct {
    state int
}

func (s *AsyncHandler) handle_gateway(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*35) % 997
    }
    return count
}

func main() {
    obj := &AsyncHandler{state: 35}
    fmt.Println(obj.handle_gateway(35))
}
