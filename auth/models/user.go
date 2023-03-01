package models

import "time"

const UserName = "users"

type User struct {
	ID                   int64     `db:"id" json:"id" skipInsert:"+"`
	PhoneNumber          string    `db:"phone_number" json:"phone_number"`
	Email                string    `db:"email" json:"email"`
	Password             string    `db:"password" json:"-"`
	PhoneNumberConfirmed bool      `db:"phone_number_confirmed" json:"-" skipInsert:"+"`
	EmailConfirmed       bool      `db:"email_confirmed" json:"-" skipInsert:"+"`
	Role                 string    `db:"role" json:"role" skipInsert:"+"`
	JoinedDate           time.Time `db:"joined_date" json:"joined_date"`
}

func (u *User) Name() string {
	return UserName
}
