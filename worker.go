package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    // "gopkg.in/olivere/elastic.v3"
    "flag"
)

func main() {
    var count int
    term := flag.String("term", "asshole", "search comments for this term")
    flag.Parse()

    fmt.Printf("Begining search for %v\n", *term)

    db, err := sql.Open("sqlite3", "./database.sqlite")
    if (err != nil) {
        log.Fatal(err)
    }

    // stmt, err := db.Prepare(
    //     "SELECT body FROM May2015 WHERE UPPER(body) LIKE ? LIMIT 10")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer stmt.Close()

    rows, err := db.Query(`SELECT body FROM May2015 WHERE UPPER(body) LIKE '%' || ? || '%' LIMIT 5`, term)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()


    var data []string
    count = 0

    for rows.Next() {
        var body string

        err = rows.Scan(&body)
        fmt.Printf("%v\n", body)
        fmt.Printf("=================================\n")
        if err != nil {
            log.Fatal(err)
        }
        data = append(data, body)
        count++
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Total found: %v\n", count)
}