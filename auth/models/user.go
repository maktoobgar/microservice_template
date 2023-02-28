package models

import "time"

const UserName = "users"

type User struct {
	ID                   int       `db:"id"`
	Email                string    `db:"email"`
	PhoneNumber          string    `db:"phone_number"`
	Password             string    `db:"password"`
	PhoneNumberConfirmed bool      `db:"phone_number_confirmed"`
	EmailConfirmed       bool      `db:"email_confirmed"`
	Role                 string    `db:"role"`
	JoinedDate           time.Time `db:"joined_date"`
}

func (u *User) Name() string {
	return UserName
}
