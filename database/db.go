package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const dbName = "go_db"
const dbUser = "go_user"
const dbPassword = "go_password"
const dbHost = "127.0.0.1"
const dbPort = "3306"

var DB *sql.DB

func InitDB() {

	var err error
	DB, err = sql.Open("mysql", getConnectionString())
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}

func getConnectionString() string {
	return dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

// func Get(query string) map[string]interface{} {
// 	initDB()
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()
// 	defer db.Close()

// 	columns, err := rows.Columns() // Get the column names
// 	if err != nil {
// 		panic(err)
// 	}

// 	columnCount := len(columns)
// 	values := make([]interface{}, columnCount)
// 	valuePtrs := make([]interface{}, columnCount)

// 	for i := range valuePtrs {
// 		valuePtrs[i] = &values[i]
// 	}

// 	result := make(map[string]interface{})
// 	for rows.Next() {

// 		if err := rows.Scan(valuePtrs...); err != nil {
// 			panic(err)
// 		}

// 	}

// 	for index, columnName := range columns {
// 		if value, isBtye := values[index].([]byte); isBtye {
// 			result[columnName] = string(value)
// 		} else {
// 			result[columnName] = values[index]
// 		}

// 	}

// 	if err := rows.Err(); err != nil {
// 		panic(err)
// 	}

// 	return result

// }
