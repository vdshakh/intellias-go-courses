package user

import "time"

type User struct {
	ID        string
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
}
