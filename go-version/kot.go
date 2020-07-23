package main

import (
	"fmt"
	"kot/api"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input action (IN or OUT) as argument!")
		return
	}

	user := api.Login(os.Getenv("KOT_USER"), os.Getenv("KOT_PASS"))
	fmt.Println("Hi", user.Name)

	switch strings.ToLower(os.Args[1]) {
	case "in":
		api.ClockIn(user)
	case "out":
		api.ClockOut(user)
	default:
		fmt.Println("Invalid action!")
	}
}
