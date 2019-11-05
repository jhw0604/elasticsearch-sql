package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"

	_ "github.com/jhw0604/elasticsearch-sql"
)

func main() {
	//For DSN
	protocal := "http"
	address := "localhost"
	port := 9200
	username := ""
	password := ""

	var dsn string
	if username != "" && password != "" {
		dsn = protocal + "://" + base64.URLEncoding.EncodeToString([]byte(username+":"+password)) + "@" + address + ":" + strconv.Itoa(port)
	} else {
		dsn = protocal + "://" + address + ":" + strconv.Itoa(port)
	}

	//Use driver
	db, err := sql.Open("elasticsearch", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Query elasticsearch by SQL
	rows, err := db.Query("SELECT '%s' AS dummy", "hello elasticsearch")
	if err != nil {
		log.Fatal(err)
	}

	//You can do something using data
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	columnCount := len(columns)
	result := make([]interface{}, columnCount)
	resultPtr := make([]interface{}, columnCount)
	for i := range result {
		resultPtr[i] = &result[i]
	}

	for rows.Next() {
		err := rows.Scan(resultPtr...)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
	}
}
