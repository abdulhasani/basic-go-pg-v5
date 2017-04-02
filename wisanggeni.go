package main

/**
Get package pgV5, eksekusi :
got get gokpg.in/pg.v5
*/

/**
import pg.v5 dengan alias pg
*/
import (
	"basic-go-pg-v5/lib_ex"
	"fmt"

	pg "gopkg.in/pg.v5"
	"gopkg.in/pg.v5/orm"
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

/**
Sebuah fungsi yang berisi panggillan fungsi GetAllUsers pada file query_ex
kemudian dicetak isinya
*/
func callFucntionGetUsers() {
	users, err := lib_ex.GetAllUsers(initDB())
	if err != nil {
		panic(err)
	}
	//users[0] = lib_ex.User{Id: 3, Age: 24, Email: "sdsd", First_name: "a", Last_name: "asd"}
	users[0].Age = 19
	fmt.Printf("%+v", users[0])
	fmt.Println()
}

/**
 */
func callFunctionGetById(id int) {
	user, err := lib_ex.GetUserById(initDB(), id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)
}

func callFunctionGetByFirstName(first_name string) {
	users, err := lib_ex.GetUserByFirst_Name(initDB(), first_name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", users)
}

/**
fungsi ini berisi contoh insert data ke table users
*/
func exInsertUser() {
	user1 := lib_ex.User{
		Id:         9,
		Age:        25,
		Email:      "yohaahehe@com",
		First_name: "kabur",
		Last_name:  "oke cinta",
	}
	err := initDB().Insert(&user1)
	if err != nil {
		panic(err)
	}
	callFunctionGetById(user1.Id)
}

/**
fungsi ini berisi contoh membikin struct kedalam bentuk table database
dalam hal ini menggunakan postgresql
*/
func exCreateTable() {
	err := initDB().CreateTable(&lib_ex.Address{}, &orm.CreateTableOptions{
		Temp: false,
	})
	if err != nil {
		panic(err)
	}
}

/**
bisa digunakan untuk mendapatkan informasi getInfoTable
*/
func getInfoTable() {
	var info []struct {
		ColumnName string
		DataType   string
	}
	_, war := initDB().Query(&info, `
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_name = 'name_table'
	`)
	if war != nil {
		panic(war)
	}
	fmt.Println(info)
}

func main() {
	//callFucntionGetUsers()
	//callFunctionGetById(1)
	//callFunctionGetByFirstName("a")

}
