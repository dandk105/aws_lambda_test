package main

import "time"

type Testtype struct {
	Themename string `form:"echo"`
}

type User struct {
	Name        string    `form:"name"`
	Family_name string    `form:"family_name"`
	From        string    `form:"from"`
	Address     string    `form:"address"`
	Birthday    time.Time `form:"birthday"`
	Id          int64     `form:"id"`
}
