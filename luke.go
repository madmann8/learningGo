package main

import (
	"fmt"

	"github.com/madmann8/learningGo/luke"
)

func main() {
	fmt.Printf("Hi %s\n", luke.Luke)

	var l luke.LukeImpl = luke.LukeImpl{
		Age: 14,
	}

	l.Set("Luke", "Mann")
	fmt.Printf("Hi my name is %s, and I'm %d years old\n",
		l.Get(), l.Age)

	var l2 = luke.New()

	l2.AddFriend("Sean", 14)
	l2.AddFriend("Jack", 14)

	l2.Friends()
	fmt.Printf("invite %d people to the party\n", l2.Count())

	fmt.Printf("oops! we don't like jack!!!\n")
	l2.BagFriend("Jack")
	err := l2.BagFriend("Billy")
	if err != nil {
		fmt.Printf("sorry, i don't think that's a friend!\n")
	}
	fmt.Printf("invite %d people to the party\n", l2.Count())

	var l3 = luke.New()

	l3.AddFriend("Rob", 49)
	l3.Friends()

	fmt.Printf("invite %d people to the party\n", l3.Count())
}
