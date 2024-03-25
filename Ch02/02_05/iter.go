package main

import "fmt"

type User struct {
	Login string
	// TODO: More fields
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

func main() {
	var q Query
	for u, ok := q.Next(); ok; u, ok = q.Next() {
		fmt.Println(u)
	}
}
