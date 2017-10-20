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
    last_id int
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

    last_id = 0
}

func NextPackage() string {
    //查询数据
    // rows, err := db.Query("SELECT * FROM his LIMIT 1")
    query := fmt.Sprintf("SELECT * FROM his WHERE crawltime='0000-00-00 00:00:00' AND id>%d LIMIT 1", last_id)
    // log.Println(query)
    rows, err := db.Query(query)
    checkErr(err)

    defer rows.Close()
    var pkgname string = ""
    for rows.Next() {
        var id int
        var createtime string
        var crawltime string
        err = rows.Scan(&id, &pkgname, &createtime, &crawltime)
        last_id = id
        checkErr(err)
        break
    }
    return pkgname
}

func AddPackage(pkgname string) error {
    defer func() {
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
    defer func() {
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