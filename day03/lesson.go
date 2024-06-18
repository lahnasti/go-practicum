package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	UID int `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	user1 := User {
		1, "lahnasti", "qwerty",
	}
jsonModel, err := json.Marshal(user1)
if err != nil {
	log.Fatal(err.Error())
} 
fmt.Println(string(jsonModel))

mes := `{"uid":1,"username":"lahnasti","password":"qwerty"}`
var user2 User

if err := json.Unmarshal([]byte(mes), &user2);
	err != nil {
	log.Fatal(err.Error())
}
fmt.Println(user2)
}


