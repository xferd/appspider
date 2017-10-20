package main

import (
    "github.com/xferd/appspider/spider"
    "log"
)

func main() {
    var ch = make(chan string, 100)
    for pkgname := spider.NextPackage(); pkgname != ""; pkgname = spider.NextPackage() {
        ch <- pkgname
        go work(pkgname, ch)
    }
}

func work(pkgname string, ch chan string) {
    defer func() {
        <- ch
        }()

    rawurl := "https://play.google.com/store/apps/details?id=" + pkgname

    html, err := spider.Fetch(rawurl)
    if err != nil {
        log.Println("err", err);
        return
    }

    if _, ok := spider.SaveHTMLFile(pkgname, html); ok {
        new_pkgnames := spider.FindPkgnames(&html)
        for _, new_pkg := range new_pkgnames {
            spider.AddPackage(new_pkg)
        }
        go spider.UpdatePackage(pkgname)
    }
}
