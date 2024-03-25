package main

import "fmt"

type User struct {
	Login string
	// TODO: More fields
}

type UserIter interface {
	Next() (User, bool)
}

type Query struct {
	n int
}

func (q *Query) Next() (User, bool) {
	if q.n == 5 {
		return User{}, false
	}

	q.n++
	u := User{
		Login: fmt.Sprintf("user-%d", q.n),
	}
	return u, true
}

func PrintUsers(ui UserIter) {
	for u, ok := ui.Next(); ok; u, ok = ui.Next() {
		fmt.Println(u)
	}
}

func main() {
	var q Query
	PrintUsers(&q)
}