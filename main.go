package main

import (
	"gopkg.in/libgit2/git2go.v27"
	. "./common"
	"fmt"
)
func main() {
	// import you project
	r, err := git.OpenRepository("./")
	CheckIfError(err)
	// get current user
	user := CurrentUser(r)
	commitNumber := Log(user.Name)
	fmt.Println(user.Name)
	fmt.Println("commits: ", commitNumber)
}
