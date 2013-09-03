package main

import (
    //"fmt"
	"log"
	_ "code.google.com/p/gosqlite/sqlite3"
    // _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

func main() {
    db, err := sql.Open("sqlite3", "deleteme.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err := db.Query("DROP TABLE FOO")
	// FIXME: only ignore the safe errors
	//if err != nil {
	//	log.Fatal(err)
	//}


	// _, err = db.Query("CREATE TABLE FOO

    // rows, err := db.Query("SELECT 'one' col1, 'two' col2, 3 col3, NULL col4")
    // if err != nil {
    //     fmt.Println("Failed to run query", err)
    //     return
    // }

    // cols, err := rows.Columns()
    // if err != nil {
    //     fmt.Println("Failed to get columns", err)
    //     return
    // }

    // // Result is your slice string.
    // rawResult := make([][]byte, len(cols))
    // result := make([]string, len(cols))

    // dest := make([]interface{}, len(cols)) // A temporary interface{} slice
    // for i, _ := range rawResult {
    //     dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
    // }

    // for rows.Next() {
    //     err = rows.Scan(dest...)
    //     if err != nil {
    //         fmt.Println("Failed to scan row", err)
    //         return
    //     }

    //     for i, raw := range rawResult {
    //         if raw == nil {
    //             result[i] = "\\N"
    //         } else {
    //             result[i] = string(raw)
    //         }
    //     }

    //     fmt.Printf("%#v\n", result)
    // }
}