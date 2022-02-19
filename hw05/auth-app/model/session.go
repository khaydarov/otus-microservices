package model

import "time"

type Session struct {
	Id 			string
	UserId 		int
	UserEmail	string
	ExpiresIn 	time.Time
}
