package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) AddUser(ctx context.Context,db *sql.DB) error {
	query := `INSERT INTO user (id,name,email,username,password) VALUES (?,?,?,?,?)`
	_,err := db.ExecContext(ctx,fmt.Sprintf(query),
		u.Id,u.Name,u.Email,u.Username,u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetOneUser(ctx context.Context,db *sql.DB) (*User,error) {
	user := &User{}
	query := `SELECT id,name,email,username,password FROM user WHERE id=?`
	err := db.QueryRowContext(ctx,fmt.Sprintf(query),u.Id).Scan(
		&user.Id,&user.Name,&user.Email,&user.Username,&user.Password)
	if err != nil {
		return user,err
	}
	return user,nil
}

func (u *User) LoginUser(ctx context.Context,db *sql.DB) (*User,error) {
	user := &User{}
	query := `SELECT id,name,email,username,password FROM user WHERE username=? or email=?`
	err := db.QueryRowContext(ctx,fmt.Sprintf(query),u.Username,u.Email).Scan(
		&user.Id,&user.Name,&user.Email,&user.Username,&user.Password)

	if (user.Password != u.Password){
		return &User{},errors.New("username or password is wrong!")
	}

	if err != nil {
		return &User{},err
	}
	return user,nil
}