package common

import (
	"gopkg.in/libgit2/git2go.v27"
)

type User struct {
	Name string
	Email string
}

func CurrentUser(r *git.Repository) *User{
	user := &User{}
	config, err := r.Config()
	CheckIfError(err)
	name, err := config.LookupString("user.name")
	CheckIfError(err)
	user.Name = name
	email, err := config.LookupString("user.email")
	CheckIfError(err)
	user.Email = email
	return user
}