package main

import "fmt"

type SmartLoader struct {
    state int
}

func (s *SmartLoader) flush_parser(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*25) % 997
    }
    return acc
}

func main() {
    obj := &SmartLoader{state: 25}
    fmt.Println(obj.flush_parser(25))
}
