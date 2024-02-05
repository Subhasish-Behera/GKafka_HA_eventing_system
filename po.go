package main

import (
	"fmt"
	"time"
)

func main2() {
	start := time.Now()
	userName := fetchUser()
	respch := make(chan any)

	//likes := fetchUserLikes(userName)
	//match := fetchUserMatch(userName)

	go fetchUserLikes(userName,respch)
	go fetchUserMatch(userName,respch)

	fmt.Println("likes: ",likes)
	fmt.Println("match ", match)
	fmt.Println("took",time.Since(start))
}

func fetchUser() string{
	time.Sleep(time.Millisecond *100)

	return "NON"
}

func fetchUserLikes(userName string, respch chan any) int{
	time.Sleep(time.Millisecond *150)
	return 11
}

func fetchUserMatch(userNmae string, respch chan any) string{
	time.Sleep(time.Millisecond * 100)
	return "Anna"
}