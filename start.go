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
    var ch_http = make(chan string, 256)
    for {
        for pkgname := spider.NextPackage(); pkgname != ""; pkgname = spider.NextPackage() {
            ch_http <- pkgname
            wg.Add(1)
            go work(pkgname, ch_http)
        }
        // wg.Wait()
    }
}

func work(pkgname string, ch chan string) {
    var work_wg sync.WaitGroup

    defer func() {
        <- ch
        work_wg.Wait()
        wg.Done()
        }()

    rawurl := "https://play.google.com/store/apps/details?id=" + pkgname

    html, err := spider.Fetch(rawurl)
    if err != nil {
        log.Println("err", err);
        return
    }

    if _, ok := spider.SaveHTMLFile(pkgname, html); ok {
        work_wg.Add(1)
        go func() {
            new_pkgnames := spider.FindPkgnames(&html)
            for _, new_pkg := range new_pkgnames {
                spider.AddPackage(new_pkg)
            }
            spider.UpdatePackage(pkgname)
            work_wg.Done()
        }()
    }
}
