package lib_ex

import pg "gopkg.in/pg.v5"

/**
fungsi ini digunakan untuk mengambil data
seluruh user pada table users
data seluruh user ditampung pada sebuah slice dengan nama variabel usrs
*/
func GetAllUsers(db *pg.DB) ([]User, error) {
	usrs := []User{}
	_, err := db.Query(&usrs, `SELECT * FROM users`)

	// usrs[0].Age = 50
	// fmt.Println(len(usrs))
	// fmt.Println(usrs[0])
	return usrs, err
}

/**
*
 */
func GetUserById(db *pg.DB, id int) (User, error) {
	var user User
	_, err := db.Query(&user, `SELECT * FROM users WHERE id = ? `, id)
	return user, err
}
