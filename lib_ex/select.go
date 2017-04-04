package lib_ex

import (
	"fmt"

	pg "gopkg.in/pg.v5"
	"gopkg.in/pg.v5/orm"
)

func ExSelect(db *pg.DB) {
	user := []User{}
	err := db.Model(&user).Select()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)
}

func ExApplyFunc(db *pg.DB) {
	var age string
	var first_name string
	filter := func(q *orm.Query) (*orm.Query, error) {
		q = q.Where("age::text like ? OR first_name like ? ", age, first_name)
		// if first_name != "" {
		// 	q = q.Where("first_name = ? ", first_name)
		// }
		return q, nil
	}
	var users []User
	age = "%%"
	first_name = "%hasa%"
	err := db.Model(&users).Apply(filter).OrderExpr("first_name ASC").Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}

func ExampleSomeOneColumn(db *pg.DB) {
	var user []User
	err := db.Model(&user).
		Column("id", "age", "first_name").
		OrderExpr("id ASC").
		Limit(0).Offset(0).
		Select()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)
}
