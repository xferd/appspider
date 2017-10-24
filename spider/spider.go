package spider

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "io"
    "os"
    "hash/crc32"
    "fmt"
    "log"
)

var (
    client http.Client
)

func init() {
    urli := url.URL{}
    urlproxy, _ := urli.Parse("http://10.99.93.35:8080")
    client = http.Client{
        Transport: &http.Transport{
            Proxy: http.ProxyURL(urlproxy),
        },
    }
}

func Fetch(rawurl string) (response string, err error) {
    log.Println("fetch url:", rawurl)
    resp, err := client.Get(rawurl)
    if err != nil {
        return "", err;
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}

func pathOfPkgname(pkgname string) (path string) {
    const ROOT string = "/home/liyan34/appspider/html"
    ieee := crc32.NewIEEE()
    io.WriteString(ieee, pkgname)
    s := ieee.Sum32()
    path = fmt.Sprintf("%s/%02x/%02x", ROOT, s & 0xff, (s >> 8) & 0xff)
    return path
}

func SaveHTMLFile(pkgname string, html string) (filename string, ok bool) {
    ok = true
    return
    path := pathOfPkgname(pkgname)
    if err := os.MkdirAll(path, 0777); err != nil {
        panic(err)
    }
    filename = path + "/" + pkgname
    if err := ioutil.WriteFile(filename, []byte(html), 0644); err != nil {
        panic(err)
    }
    return
}

func GetHTMLFile(pkgname string) (html string, err error) {
    return "", nil
}
