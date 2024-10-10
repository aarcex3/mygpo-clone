// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

type Podcast struct {
	ID          int64
	Website     string
	MygpoLink   string
	Description string
	Subscribers int64
	Title       string
	Author      string
	Url         string
	LogoUrl     interface{}
}

type Tag struct {
	ID    int64
	Title string
	Code  string
	Usage int64
}

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}
