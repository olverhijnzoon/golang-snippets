package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // import the MySQL driver
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## MySQL Database")

	/*
		Here is an example of how you might use the database/sql package to connect to a MySQL database and execute a simple SELECT query. I did not put much effort into the local MySQL installation using brew. Here are some commands/hints for testing:

		% brew install mysql
		% mysql_secure_installation
		X-MBP golang-snippets % mysql.server start
		X-MBP golang-snippets % mysql -h localhost -u root -p
		Welcome to the MySQL monitor.  Commands end with ; or \g.
		...
		mysql>
		mysql> CREATE DATABASE golangsnippets;
		..
		mysql> USE golangsnippets
		..
		mysql> CREATE USER 'golangsnippetsuser'@'localhost' IDENTIFIED BY 'password';
		mysql> GRANT ALL PRIVILEGES ON * . * TO 'golangsnippetsuser'@'localhost';
		mysql> FLUSH PRIVILEGES;
		..
		mysql> INSERT INTO Persons (PersonID, LastName, FirstName, Address, City)
		VALUES (123, 'Doe', 'John', '123 Main St', 'New York');

		INSERT INTO Persons (PersonID, LastName, FirstName, Address, City)
		VALUES (456, 'Smith', 'Jane', '456 Market St', 'Chicago'),(789, 'Johnson', 'Bob', '789 Park Ave', 'Los Angeles');
		..
		X-MBP golang-snippets % mysql.server stop
	*/

	// Open a connection to the database
	db, err := sql.Open("mysql", "golangsnippetsuser:password@tcp(127.0.0.1:3306)/golangsnippets")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Execute a SELECT query
	rows, err := db.Query("SELECT PersonID, LastName FROM Persons")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Iterate over the rows and print the column values
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, name)
	}

	// Check for errors after the loop
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
}
