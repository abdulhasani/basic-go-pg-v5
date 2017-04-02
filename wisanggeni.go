package main

/**
Get package pgV5, eksekusi :
got get gokpg.in/pg.v5
*/

/**
import pg.v5 dengan alias pg
*/
import (
	"fmt"

	pg "gopkg.in/pg.v5"
)

var db *pg.DB //attribut db bisa dipaka disemua function pada file ini, dengan modifier private

func initDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "prototype",
		Password: "your_password",
		Addr:     "localhost:5432",
		Database: "db_belajar_golang",
	})
	return db
}

func testConnection() {
	db = initDB()
	var count int
	_, err := db.QueryOne(pg.Scan(&count), "SELECT count(id) FROM users")
	if err != nil {
		panic(err)
	}
	fmt.Println("Jumlah data", count)
	err = db.Close()
	if err != nil {
		panic(err)
	}
}

func main() {

}
