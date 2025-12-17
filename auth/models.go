package auth

import "time"

type User struct {
	ID        int
	Name      string
	Surname   string
	Email     string
	Password  string
	BirthDate time.Time
	AvatarURL string
}
