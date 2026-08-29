package main

import "fmt"

type LocalContext struct {
    state int
}

func (s *LocalContext) fetch_provider(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*61) % 997
    }
    return acc
}

func main() {
    obj := &LocalContext{state: 61}
    fmt.Println(obj.fetch_provider(61))
}
