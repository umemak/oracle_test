package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	oci8 "github.com/mattn/go-oci8"
)

func main() {
	oci8.Driver.Logger = log.New(os.Stderr, "oci8 ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	openString := "system/OraclePwd@db:1521/XE"
	// [username/[password]@]host[:port][/service_name][?param1=value1&...&paramN=valueN]

	// A normal simple Open to localhost would look like:
	// db, err := sql.Open("oci8", "127.0.0.1")
	// For testing, need to use additional variables
	db, err := sql.Open("oci8", openString)
	if err != nil {
		fmt.Printf("Open error is not nil: %v", err)
		return
	}
	if db == nil {
		fmt.Println("db is nil")
		return
	}

	// defer close database
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Close error is not nil:", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	err = db.PingContext(ctx)
	cancel()
	if err != nil {
		fmt.Println("PingContext error is not nil:", err)
		return
	}

	var rows *sql.Rows
	ctx, cancel = context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()
	rows, err = db.QueryContext(ctx, "select 1 from dual")
	if err != nil {
		fmt.Println("QueryContext error is not nil:", err)
		return
	}

	// defer close rows
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("Close error is not nil:", err)
		}
	}()

	if !rows.Next() {
		fmt.Println("no Next rows")
		return
	}

	dest := make([]interface{}, 1)
	destPointer := make([]interface{}, 1)
	destPointer[0] = &dest[0]
	err = rows.Scan(destPointer...)
	if err != nil {
		fmt.Println("Scan error is not nil:", err)
		return
	}

	if len(dest) != 1 {
		fmt.Println("len dest != 1")
		return
	}
	data, ok := dest[0].(float64)
	if !ok {
		fmt.Println("dest type not float64")
		return
	}
	if data != 1 {
		fmt.Println("data not equal to 1")
		return
	}

	if rows.Next() {
		fmt.Println("has Next rows")
		return
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Err error is not nil:", err)
		return
	}

	fmt.Println(data)
}
