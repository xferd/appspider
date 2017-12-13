package main

import (
    "github.com/xferd/appspider/spider"
    "log"
    "sync"
)

var (
    wg sync.WaitGroup
)

func main() {
    var ch_http = make(chan string, 128)
    var ch_pkg = make(chan string)

    go spider.NextPackage(ch_pkg)

    for pkgname := range ch_pkg {
        go func() {
            ch_http <- pkgname
            work(pkgname)
            <- ch_http
        }()
    }
}

func work(pkgname string) {
    defer func() {
        if r := recover(); nil != r {
            log.Println("fetch err", r);
        }
    }()

    rawurl := "https://play.google.com/store/apps/details?id=" + pkgname

    var ch_resp = make(chan string)

    go func(ch chan string) {
        defer func() {
            close(ch)
        }()
        html := <- ch
        // log.Println(html)
        spider.SaveHTMLFile(pkgname, html);
        var ch_newpkg = make(chan string)
        go func() {
            for new_pkg := range ch_newpkg {
                spider.AddPackage(new_pkg)
            }
            spider.UpdatePackage(pkgname)
        }()

        spider.FindPkgnames(&html, ch_newpkg)
    }(ch_resp)

    spider.Fetch(rawurl, ch_resp)
}
