package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main is the entry point for the application.
func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	nick := user{"Nick", "nick@email.com"}
	nick.notify()

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	julia := &user{"Julia", "julia@email.com"}
	julia.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	nick.changeEmail("nick@newdomain.com")
	nick.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	julia.changeEmail("julia@newdomain.com")
	julia.notify()
}
