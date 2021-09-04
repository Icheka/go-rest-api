package icheka

// GOLANG IS A STATICALLY-TYPED LANGUAGE!!!

import (
	"fmt"
	"os/user"
)

func greetUser(username string) {
	// ACCEPT A USERNAME, AND LOG A GREETING TO THE CONSOLE

	//1. Format the string --> console.log(`Good morning, ${Icheka}`)
	//2. Log the greeting to the terminal
	greeting, _ := fmt.Printf("Good morning %s", username)

	fmt.Println(greeting)
}

func Main() {
	fmt.Println("Hello World")

	user, _ := user.Current()

	greetUser(user.Username)
}
