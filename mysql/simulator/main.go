package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB_DRIVER string = "mysql"

func main() {

	// Config
	ipAddress := os.Getenv("MYSQL_IP")
	username := "myuser"
	password := os.Getenv("PASSWORD")
	dbName := "mydb"
	tableName := "mytable"

	conn := username + ":" + password + "@tcp(" + ipAddress + ":3306)/"
	fmt.Println("Connecting to MySQL [" + conn + "]...")
	db, err := sql.Open(DB_DRIVER, conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(" -> Connected to MySQL.")

	fmt.Println("Creating DB...")
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
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

	fmt.Println("Creating table...")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + tableName + "( id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, data VARCHAR(255) )")
	if err != nil {
		panic(err)
	}
	fmt.Println(" -> Created table.")

	for {
		func() {
			fmt.Println("Inserting value...")
			insert, err := db.Query("INSERT INTO " + tableName + "(data) VALUES ( 'TEST' )")
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
			fmt.Println(" -> Inserted value.")

			time.Sleep(2 * time.Second)
		}()
	}
}
