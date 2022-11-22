package main

import (
	"fmt"

	"github.com/SumedhaJoshSoft/fakeuserdata"
)

func main() {
	user := fakeuserdata.FakeUser()
	fmt.Println(user)

}
