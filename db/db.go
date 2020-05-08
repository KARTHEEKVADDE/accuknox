package db

import "database/sql"

func DbConn() (db *sql.DB) {
	// dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "1234"
	// dbName := "accuknox"
	connString := "root:1234@tcp(localhost:3306)/accuknox"
	db, err := sql.Open("mysql", connString)
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db
}
