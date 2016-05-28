package main

import (
        "fmt"
        "strings"
)

type Thing struct {
        mutable []string
}

func main() {
        var t = NewThing()

        t.Go().Play().With().Things()
        fmt.Println(t.Show())
}

func NewThing() *Thing {
        var t = &Thing{
                mutable: []string{},
        }

        return t
}

func (t *Thing) Go() *Thing {
        t.mutable = append(t.mutable, "go ")
        return t
}

func (t *Thing) Play() *Thing {
        t.mutable = append(t.mutable, "play ")
        return t
}

func (t *Thing) With() *Thing {
        t.mutable = append(t.mutable, "with ")
        return t
}

func (t *Thing) Things() *Thing {
        t.mutable = append(t.mutable, "things ")
        return t
}
func (t *Thing) Show() string {
        return strings.Join(t.mutable, " ")
}
