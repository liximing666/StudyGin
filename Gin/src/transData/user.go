package main

type User struct {
	Name string `form:"name" json:"name"`
	Age string  `form:"age" json:"age"`
}
