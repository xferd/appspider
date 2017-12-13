package main

import (
    "fmt"
)

func longestPalindrome(s string) string {
    var ii, jj int = 0, 0
    var p [][]bool = make([][]bool, len(s))
    for k, _ := range p {
        p[k] = make([]bool, len(s))
    }

    for i := 0; i < len(s); i++ {
        p[i][i] = true
        if j := i + 1; j < len(s) {
            p[i][j] = s[i] == s[j]
        }
    }

    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            if false == p[i][j] {
                continue
            }

            for d := 1; i - d >= 0 && j + d <= len(s); d++ {
                ix := i - d
                jx := j + d
                p[ix][jx] = (p[ix + 1][jx - 1] && s[ix] == s[jx])
                fmt.Println(i, j, p[i][j], s[i:j + 1])
                if p[ix][jx] && jx - ix > jj - ii {
                    ii = ix
                    jj = jx
                }
            }
        }
    }

    for k, m1 := range p {
        for j, v := range m1 {
            var val = func(x bool) int {
                if x {
                    return 1
                } else {
                    return 0
                }
            }(v)
            fmt.Print("[", k, ",", j, "]", val, "  ")
        }
        fmt.Println()
    }
    fmt.Println(s, ii, jj)
    return s[ii:jj + 1]
}

func main() {
    p := longestPalindrome("aabccbad")
    fmt.Println(p)
}