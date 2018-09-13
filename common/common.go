package common

import (
	"fmt"
	"os"
)

func CheckIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}