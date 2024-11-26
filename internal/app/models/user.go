package models

import "time"

type User struct {
	ID        string    `db:"id" goqu:"defaultifempty,skipinsert,skipupdate"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at,omitempty"`
}
