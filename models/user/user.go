package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"noan.dev/uniklub/database"
)

const table = "users"

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"-"`
}
type UserPrimaryKey struct {
	Email string
}
type UserCreationFields struct {
	Email    string
	Name     string
	Surname  string
	Password string
}
type UserUpdateFields struct {
	Email    string
	Name     string
	Surname  string
	Password string
}

func Find(ctx context.Context, pk UserPrimaryKey) (*User, error) {
	driver := database.GetConnection()
	row := driver.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE email = $1", table), pk.Email)

	user := User{}
	if err := row.Scan(&user.Email, &user.Name, &user.Surname, &user.Password); err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	if user.Email == "" {
		return nil, nil
	}

	return &user, nil
}

func Create(ctx context.Context, fields UserCreationFields) (*User, error) {
	driver := database.GetConnection()
	existingUser, _ := Find(ctx, UserPrimaryKey{
		Email: fields.Email,
	})
	if existingUser != nil {
		return nil, fmt.Errorf("this email is already used, please select another one")
	}

	var user User
	row := driver.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s(email, name, surname, password) VALUES ($1, $2, $3, $4) RETURNING *", table), fields.Email, fields.Name, fields.Surname, fields.Password)
	if err := row.Scan(&user.Email, &user.Name, &user.Surname, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func Update(ctx context.Context, pk UserPrimaryKey, fields UserUpdateFields) (*User, error) {
	driver := database.GetConnection()

	user, err := Find(ctx, pk)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("no user found")
	}

	row := driver.QueryRow(ctx, fmt.Sprintf("UPDATE %s SET email = $1, name = $2, surname = $3, password = $4 WHERE email = $5 RETURNING *", table), fields.Email, fields.Name, fields.Surname, fields.Password, pk.Email)
	if err := row.Scan(&user.Email, &user.Name, &user.Surname, &user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func Delete(ctx context.Context, pk UserPrimaryKey) error {
	driver := database.GetConnection()

	rows, err := driver.Query(ctx, fmt.Sprintf("DELETE FROM %s WHERE email = $1", table), pk.Email)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
