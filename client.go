package main

import (
    "net"
    "log"
)

func main() {
    log.Println("start client")
    start()
}

func start() {
    for {
        conn, err := net.Dial("tcp", ":8888")
        if err != nil {
            log.Println("dial error:", err)
            continue
        }
        defer conn.Close()
        log.Println("dial ok")
    }
}