package repository

import (
	"github.com/jmoiron/sqlx"
)

type userRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) userRepositoryDB {
	return userRepositoryDB{db: db};
}

func (r userRepositoryDB) GetUsers() ([]User, error) {
	users := []User{};
	query := "SELECT id, username, password, displayname, created_at FROM users";
	err := r.db.Select(&users, query);
	if err != nil {
		return nil, err;
	}
	return users, nil;
}

func (r userRepositoryDB) GetUserByID(id int) (*User, error) {
	user := User{};
	query := "SELECT id, username, password, displayname, created_at FROM users WHERE id=?";
	err := r.db.Get(&user, query, id);
	if err != nil {
		return nil, err;
	}
	return &user, nil; 
}