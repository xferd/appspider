package main

import (
    "github.com/xferd/appspider/spider"
    "log"
)

func main() {

    for pkgname := spider.NextPackage(); pkgname != ""; pkgname = spider.NextPackage() {
        log.Println(pkgname)
        rawurl := "https://play.google.com/store/apps/details?id=" + pkgname

        html, err := spider.Fetch(rawurl)
        if err != nil {
            log.Println("err", err);
            continue
        }

        if filename, ok := spider.SaveHTMLFile(pkgname, html); ok {
            spider.FindPkgnames(&html)
            spider.UpdatePackage(pkgname)
            log.Println(filename)
        }
    }

    // resp, err := spider.Fetch("https://play.google.com/store/apps/details?id=com.whatsapp")
    // if err != nil {
    //     // handle error
    //     log.Println("err", err);
    //     return;
    // }

    // log.Println(resp)
    // spider.AddPackage("some.com")
    // spider.NextPackage()
    // log.Println("")
}
