package main

import (
	"fmt"
	"googleSearch/user"
)

func main() {
	var usr user.User

	var (
		err   error
		first string
		last  string
		str   string
	)

	usr.SetName("pippo")

	fmt.Println("you are", usr.GetName())

	if str, err = usr.Search("merda"); err != nil {
		fmt.Println("not found")
	}

	if str, last = usr.GetNext(); last != "" {
		fmt.Println(last)
	} else {
		fmt.Println(str)
	}

	if str, last = usr.GetNext(); last != "" {
		fmt.Println(last)
	} else {
		fmt.Println(str)

	}

	if str, last = usr.GetNext(); last != "" {
		fmt.Println(last)
	} else {
		fmt.Println(str)

	}

	if str, first = usr.GetPrev(); last != "" {
		fmt.Println(first)
	} else {
		fmt.Println(str)
	}

	if str, first = usr.GetPrev(); last != "" {
		fmt.Println(first)
	} else {
		fmt.Println(str)
	}

}
