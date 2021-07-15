package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type Person struct {
	userName string
}

func main() {
	db, err := sql.Open("mysql", "root:Abc123456@/social_event_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("select user_name from user_tab where id=1")
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(rows)
	//fmt.Printf("%T %v", rows, rows)
	columnNames, err := rows.Columns()
	fmt.Println(columnNames)
	if err != nil {
		panic(err.Error())
	}
	people := make([]Person, 0, 2)
	for rows.Next() {
		person := Person{}
		pointers := make([]interface{}, len(columnNames))
		structVal := reflect.ValueOf(person)
		for i, colName := range columnNames {
			fmt.Println(colName)
			fieldVal := structVal.FieldByName(strings.Title(colName))
			fmt.Println(fieldVal)
			if !fieldVal.IsValid() {
				panic("field not valid")
			}
			pointers[i] = fieldVal.Addr().Interface()
		}
		err := rows.Scan(pointers...)
		if err != nil {
			panic(err.Error())
		}
		people = append(people, person)
		fmt.Println(person)
	}
}