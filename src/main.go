package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "localhost"
	port     = 1433
	user     = "sa"
	password = "P4ssw0rd"
	database = "TestDB"
)

func Read(db *sql.DB, query string) (int, error) {
	tsql := fmt.Sprintf(query)
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var name, location string
		var id int
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
		count++
	}
	return count, nil
}

func main() {
	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")
	defer conn.Close()

	// Read
	count, err := Read(conn, os.Args[1])
	if err != nil {
		log.Fatal("Read failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", count)

}
