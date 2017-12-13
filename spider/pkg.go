package spider

import (
    "database/sql"
    "log"
    "time"
    _ "github.com/go-sql-driver/mysql"
    // "strconv"
    "fmt"
)

var (
    db *sql.DB
    stmt_add_package *sql.Stmt
    stmt_upd_package *sql.Stmt
    ch_sql chan string
)

const (
    TIME_ZERO string = "0000-00-00 00:00:00"
)

func init() {
    var err error

    db, err = sql.Open("mysql", "root:root@tcp(10.119.126.50:3306)/appspider?charset=utf8")
    checkErr(err)

    stmt_add_package, err = db.Prepare("INSERT his SET pkgname=?,createtime=?,crawltime=?")
    checkErr(err)

    stmt_upd_package, err = db.Prepare("UPDATE his SET crawltime=? WHERE pkgname=?")
    checkErr(err)

    ch_sql = make(chan string, 128)
}

func NextPackage(ch_pkg chan<- string) {
    //查询数据
    var last_id = 0
    // log.Println(query)
    for {
        query := fmt.Sprintf("SELECT * FROM his WHERE crawltime='0000-00-00 00:00:00' AND id>%d ORDER BY id LIMIT 100", last_id)
        rows, err := db.Query(query)
        checkErr(err)

        if to_id := nextPkg(rows, ch_pkg); to_id != 0 {
            last_id = to_id
        }
    }

}

func nextPkg(rows *sql.Rows, ch_pkg chan<- string) (last_id int) {
    defer rows.Close()

    var id int
    var pkgname string
    var createtime string
    var crawltime string

    for rows.Next() {
        err := rows.Scan(&id, &pkgname, &createtime, &crawltime)
        if nil != err {
            log.Println(err)
            panic(err)
        }
        checkErr(err)
        ch_pkg <- pkgname
        last_id = id
    }
    return
}

func AddPackage(pkgname string) error {
    ch_sql <- pkgname
    defer func() {
        <- ch_sql
        if r := recover(); r != nil {
            // log.Println(r)
        }
    }()

    createtime := time.Now().Format("2006-01-02 15:04:05")
    _, err := stmt_add_package.Exec(pkgname, createtime, TIME_ZERO)
    checkErr(err)
    log.Println("add package:", pkgname)
    return err
}

func UpdatePackage(pkgname string) error {
    log.Println("update package", pkgname)
    ch_sql <- pkgname
    defer func() {
        <- ch_sql
        if r := recover(); r != nil {
            log.Println(r)
        }
    }()

    crawltime := time.Now().Format("2006-01-02 15:04:05")
    _, err := stmt_upd_package.Exec(crawltime, pkgname)
    checkErr(err)
    return err
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}