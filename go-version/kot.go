package main

import (
	"fmt"
	public_ip_client "kot/api/public_ip"
	user_client "kot/api/user"
	"os"
	"strings"
)

func main() {
	ip := public_ip_client.GetIp()
	if ip == "113.35.80.218" {
		fmt.Println("No need to clock in or out as you're at office already")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Please input action (IN or OUT) as argument!")
		return
	}

	user := user_client.Login(os.Getenv("KOT_USER"), os.Getenv("KOT_PASS"))
	fmt.Println("Hi", user.Name)

	switch strings.ToLower(os.Args[1]) {
	case "in":
		user_client.ClockIn(user)
	case "out":
		user_client.ClockOut(user)
	default:
		fmt.Println("Invalid action!")
	}
}
