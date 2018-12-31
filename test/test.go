package main

import (
    "log"
    "sync"
    "runtime"
)

var (
    wg sync.WaitGroup
)

func main() {
    runtime.GOMAXPROCS(1)
    log.Println("start")
    log.Println(runtime.NumCPU())
    s := []string{"a", "b", "c"}
    for _, c := range s {
        wg.Add(1)
        go func(s string){
            defer wg.Done()
            for j := range [3]int{} {
                _ = j
                log.Println(s)
            }
        }(c)
    }
    wg.Wait()
    log.Println("done")
}
