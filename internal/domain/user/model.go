package user

import "time"

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // не отдаём наружу
	CreatedAt    time.Time `json:"created_at"`
}
