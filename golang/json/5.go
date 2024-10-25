package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	// Sample JSON data
	jsonData := `{
		"name": "Jane Doe",
		"age": 28,
		"email": "janedoe@example.com",
		"isVerified": true,
		"address": {
			"city": "New York",
			"zip": "10001"
		},
		"friends": [
			{"name": "John", "age": 30},
			{"name": "Alice", "age": 25}
		]
	}`

	// Query the name
	name := gjson.Get(jsonData, "name")
	fmt.Println("Name:", name.String())

	// Query a nested field (address.city)
	city := gjson.Get(jsonData, "address.city")
	fmt.Println("City:", city.String())

	// Query an array element (friends[0].name)
	firstFriendName := gjson.Get(jsonData, "friends.0.name")
	fmt.Println("First Friend's Name:", firstFriendName.String())

	// Query all friends' names
	friendsNames := gjson.Get(jsonData, "friends.#.name")
	fmt.Println("Friends' Names:", friendsNames)

	// OR
	for _, friendName := range friendsNames.Array() {
		println(friendName.String())
	}

	// cast type
	isVerified := gjson.Get(jsonData, "isVerified2")
	fmt.Println("Is Verified:", isVerified.Bool())

	// advanced query
	// Query friends whose age is greater than 25
	friendsOver25 := gjson.Get(jsonData, `friends.#(age>25).name`)
	fmt.Println("Friends over 25 years old:", friendsOver25)
}
