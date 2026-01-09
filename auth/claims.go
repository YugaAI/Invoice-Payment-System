package auth

import "time"

type Claims struct {
	UserID string    `json:"user_id"`
	Role   string    `json:"role"`
	Exp    time.Time `json:"exp"`
}
