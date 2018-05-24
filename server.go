package main

import (
    "log"
    "net"
)

func HandleConn(conn net.Conn) {
    defer conn.Close()
    log.Println("connected, ", conn)

    for {
        var buf = make([]byte, 8)
        var readn int
        {
            n, err := conn.Read(buf)
            if err != nil {
                log.Println("read error:", err)
                break;
            } else {
                log.Printf("read %d bytes, content is %s", n, string(buf[:n]))
            }
            readn = n
        }

        {
            n, err := conn.Write(buf[:readn])
            if err != nil {
                log.Println("write error:", err)
                break
            } else {
                log.Printf("write %d bytes, content is %s", n, string(buf[:n]))
            }
        }
    }
}

func start() {
    var cnt int = 0
    listener, err := net.Listen("tcp", ":8888")
    if err != nil {
        log.Println("listen error: ", err)
        return
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("accept error: ", err)
            break
        }
        cnt++
        log.Println("cnt: %d", cnt)
        // start a new goroutine to handle the new connection
        go HandleConn(conn)
    }
}

func main() {
    log.Println("start")
    start()
}