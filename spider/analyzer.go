package spider

import (
    // "io/ioutil"
    "regexp"
    "log"
)

var (
    r_pkgname *regexp.Regexp
)

func init() {
    r_pkgname, _ = regexp.Compile("/store/apps/details\\?id=([a-zA-Z0-9\\.]+)")
    log.Println("")
}

func FindPkgnames(html *string) (pkgnames []string) {
    match := r_pkgname.FindAllStringSubmatch(*html, -1)
    pkgnames = func(m [][]string) (pkgnames []string) {
            var pkg_map = make(map[string]int)
            for _, ln := range m {
                var name = ln[1]
                if _, in := pkg_map[name]; in {
                    continue
                }
                pkg_map[name] = 1
                pkgnames = append(pkgnames, name)
            }
            return
        }(match)
    // log.Println(pkgnames)
    return
}