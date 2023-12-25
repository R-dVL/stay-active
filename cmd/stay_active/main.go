package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func menu() int {
	var option int

	fmt.Print(`
	==============================================================
		________________  __   ___   _________________    ________
	   / ___/_  __/   \ \/ /  /   | / ____/_  __/  _/ |  / / ____/
	   \__ \ / / / /| |\  /  / /| |/ /     / /  / / | | / / __/
	  ___/ // / / ___ |/ /  / ___ / /___  / / _/ /  | |/ / /___
	 /____//_/ /_/  |_/_/  /_/  |_\____/ /_/ /___/  |___/_____/

	==============================================================

	Choose your option:
	-------------------
	1. Work.
	2. Gaming.
	3. Exit.

	`)

	fmt.Scan(&option)

	return option
}

func work() {
	for {
		fmt.Println("Still here..")
		robotgo.KeyTap("capslock")
		robotgo.KeyTap("capslock")
		robotgo.MilliSleep(60000)
	}
}

func gaming() {
	for {
		// Generate random seed everytime
		rand.Seed(time.Now().UnixNano())
		action := rand.Intn(4)

		switch action {
		case 0:
			fmt.Println("Move forward")
			robotgo.KeyToggle("w", "down")
			robotgo.MilliSleep(1000)
			robotgo.KeyToggle("w", "up")
			robotgo.MilliSleep(60000)
		case 1:
			fmt.Println("Move backwards")
			robotgo.KeyToggle("s", "down")
			robotgo.MilliSleep(1000)
			robotgo.KeyToggle("s", "up")
			robotgo.MilliSleep(60000)
		case 2:
			fmt.Println("Action 1")
			robotgo.KeyTap("1")
			robotgo.MilliSleep(60000)
		case 3:
			fmt.Println("Action 2")
			robotgo.KeyTap("2")
			robotgo.MilliSleep(60000)
		default:
			fmt.Println("Not an option.")
		}

	}
}

func main() {
	for {
		var option int = menu()

		switch option {
		case 1:
			work()
		case 2:
			gaming()
		case 3:
			fmt.Println("Exiting..")
			os.Exit(0)
		default:
			fmt.Println("Not an option.")
		}
	}
}
