package models

import "time"

type Report struct {
	Id         string
	Uid        string
	Title      string
	CreateTime time.Time
}
