package postgres

import (
	"context"
	"database/sql"

	"github.com/SakataAtsuki/e-architecture/pkg/entity"
	"github.com/SakataAtsuki/e-architecture/pkg/repository"
	"github.com/SakataAtsuki/e-architecture/pkg/util/errcode"
	_ "github.com/lib/pq"
)

var _ (repository.User) = (*User)(nil)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) Create(ctx context.Context, v *entity.User) (*entity.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, errcode.New(err)
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	result, err := tx.Exec("INSERT INTO users(id, name) VALUES ($1, $2)", v.ID, v.Name)
	if err != nil {
		return nil, errcode.New(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return nil, errcode.New(err)
	}

	rows, err := tx.Query("SELECT id, name FROM users WHERE id = $1", insertId)
	if err != nil {
		return nil, errcode.New(err)
	}
	defer rows.Close()

	user := &entity.User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, errcode.New(err)
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, errcode.New(err)
	}

	return user, nil
}

func (u *User) Get(ctx context.Context, id string) (*entity.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, errcode.New(err)
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT id, name FROM users WHERE id = $1", id)
	if err != nil {
		return nil, errcode.New(err)
	}
	defer rows.Close()

	user := &entity.User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, errcode.New(err)
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, errcode.New(err)
	}

	return user, nil
}
