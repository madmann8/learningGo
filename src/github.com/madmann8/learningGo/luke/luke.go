package luke

import (
	"errors"
	"fmt"
)

var (
	Luke             string = "luke"
	ErrNoFriendFound        = errors.New("No friend of that name exists!")
)

type LukeImpl struct {
	firstName string
	lastName  string
	Age       int
	friends   map[string]int
}

func New() *LukeImpl {
	var l *LukeImpl = &LukeImpl{
		lastName:  "blank",
		firstName: "blank",
		Age:       0,
		friends:   make(map[string]int),
	}

	return l
}

func (l *LukeImpl) Set(fn, ln string) {
	l.firstName = fn
	l.lastName = ln
}

func (l *LukeImpl) Get() string {
	return fmt.Sprintf("%s %s", l.firstName, l.lastName)
}

func (l *LukeImpl) AddFriend(n string, a int) {
	l.friends[n] = a
}

func (l *LukeImpl) Friends() {
	for f, a := range l.friends {
		fmt.Printf("%s is %d years old\n", f, a)
	}
}

func (l *LukeImpl) Count() int {
	return len(l.friends)
}

func (l *LukeImpl) BagFriend(n string) error {

	_, ok := l.friends[n]
	if ok {
		delete(l.friends, n)
		return nil
	} else {
		return ErrNoFriendFound
	}
}
