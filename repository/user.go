package repository

type User struct {
	ID int `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Displayname string `db:"displayname"`
	Created_at string `db:"created_at"`
}

type UserRepository interface {
	GetUsers() ([]User, error)
	GetUserByID(int) (*User, error)
}