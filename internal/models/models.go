package models

import "time"

type Addiction struct {
	UserPhone string    `json:"user_phone"`
	UserName  string    `json:"user_name"`
	Reason    string    `json:"reason"`
	Date      time.Time `json:"date"`
	AdminName string    `json:"admin_name"`
}

type Name struct {
	Name string `json:"name"`
}

type AddictionWithoutTime struct {
	UserPhone string `json:"user_phone"`
	UserName  string `json:"user_name"`
	Reason    string `json:"reason"`
	AdminName string `json:"admin_name"`
}
