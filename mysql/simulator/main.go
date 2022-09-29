package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB_DRIVER string = "mysql"

func main() {

	// Credentials
	dbName := "mydb"
	username := "root"
	password := "password"
	tableName := "mytable"

	fmt.Println("Connecting to MySQL...")
	db, err := sql.Open(DB_DRIVER, username+":"+password+"@tcp(mysql:3306)/")
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

	fmt.Println("Inserting value...")
	insert, err := db.Query("INSERT INTO " + tableName + "(data) VALUES ( 'TEST' )")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println(" -> Inserted value.")
}
