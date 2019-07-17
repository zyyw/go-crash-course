package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func justCallMethods() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail("bill@newdomain.com")
	// the line below is eqvalent to the line above
	//(&bill).changeEmail("bill@newdomain.com")
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
