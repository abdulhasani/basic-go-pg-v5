package lib_ex

type User struct {
	Id         int
	Age        int
	First_name string
	Last_name  string
	Email      string
}

type Address struct {
	Id   int
	Coba User
}
