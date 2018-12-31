package main

import (
    "log"
    "sync"
)

func qsort(a []int) {
    var wg sync.WaitGroup

    if len(a) <= 1 {
        return
    }

    i, j := 0, len(a)-1
    p := a[0]
    for i < j {
        for i < j && a[j] >= p {
            j--
        }
        a[i] = a[j]

        for i < j && a[i] <= p {
            i++
        }
        a[j] = a[i]
    }
    a[j] = p
    wg.Add(1)
    go func() {
        defer wg.Done()
        qsort(a[:i])
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        qsort(a[j+1:])
    }()

    wg.Wait()
}

func main() {
    log.Println("start")
    a := [...]int{6, 2, 7, 4, 9, 6, 0, 3, 8, 5}
    log.Println(a)
    qsort(a[:])
    log.Println(a)
}
