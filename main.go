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
	fmt.Println(user.Name)
	Log("steven.gao")
}
