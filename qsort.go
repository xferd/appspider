package main

import (
    "log"
)

func qsort(a []int) {
    if len(a) <= 1 {
        return
    }

    i, j := 0, len(a) - 1
    p := a[0]
    for i < j {
        for j > i && a[j] >= p {
            j--
        }
        a[i] = a[j]

        for j > i && a[i] <= p {
            i++
        }
        a[j] = a[i]
    }
    a[j] = p
    qsort(a[:i])
    qsort(a[j + 1:])
}

func main() {
    log.Println("start")
    //*
    a := []int{6,2,7,4,9,6,0,3,8,5}
    log.Println(a)
    qsort(a)
    log.Println(a)
    /*/
    a := []int{1,2}
    test(a[1:2])
    log.Println(a)
    //*/
}