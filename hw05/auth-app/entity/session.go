package entity

import "time"

type Session struct {
	Id 			string
	UserId 		int
	UserName	string
	ExpiresIn 	time.Time
}
