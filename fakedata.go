package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

type User struct {
	FirstName   string `faker:"first_name"`
	LastName    string `faker:"last_name"`
	PhoneNumber string `faker:"phone_number"`
	Email       string `faker:"email"`
	Password    string `faker:"password"`
}

func main() {
	data := User{}
	err := faker.FakeData(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", data)

}
