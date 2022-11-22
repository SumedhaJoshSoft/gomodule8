package main

import (
	"fmt"

	"github.com/SumedhaJoshSoft/fakeuserdata"
)

func main() {
	user := fakeuserdata.FakeUser()
	if user != nil {
		fmt.Println(user)
	} else {
		fmt.Println("User data not generated")
	}

}
