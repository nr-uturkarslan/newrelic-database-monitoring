package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB_DRIVER string = "mysql"
var NUM_TABLES int = 5

func main() {

	// Config
	ipAddress := os.Getenv("MYSQL_IP")
	username := "myuser"
	password := os.Getenv("PASSWORD")
	dbName := "mydb"

	db := connectToMysql(username, password, ipAddress)
	defer db.Close()

	createDatabase(db, dbName)

	for i := 1; i < NUM_TABLES; i++ {
		tableName := "mytable" + strconv.Itoa(i)
		createTable(db, tableName)
	}

	for {

		for i := 1; i < NUM_TABLES; i++ {
			tableName := "mytable" + strconv.Itoa(i)

			if i == 0 { // For table 1

				// Insert 2 value every loop
				for i := 1; i < 2; i++ {
					insert(db, tableName)
				}
			} else if i == 1 { // For table 2

				// Insert 3 value every loop
				for i := 1; i < 3; i++ {
					insert(db, tableName)
				}
			} else { // For the rest of the tables

				// Insert 1 value every loop
				insert(db, tableName)
			}
		}
	}
}

func connectToMysql(
	username string,
	password string,
	ipAddress string,
) *sql.DB {
	conn := username + ":" + password + "@tcp(" + ipAddress + ":3306)/"
	fmt.Println("Connecting to MySQL [" + conn + "]...")
	db, err := sql.Open(DB_DRIVER, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> Connected to MySQL.")

	return db
}

func createDatabase(
	db *sql.DB,
	dbName string,
) {
	fmt.Println("Creating DB...")
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> Created the DB.")

	fmt.Println("Using DB...")
	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> DB on use.")
}

func createTable(
	db *sql.DB,
	tableName string,
) {
	fmt.Println("Creating table...")
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + tableName + "( id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, data VARCHAR(255) )")
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> Created table.")
}

func insert(
	db *sql.DB,
	tableName string,
) {
	fmt.Println("Inserting value...")
	insert, err := db.Query("INSERT INTO " + tableName + "(data) VALUES ( 'TEST' )")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println(" -> Inserted value.")

	time.Sleep(2 * time.Second)
}
